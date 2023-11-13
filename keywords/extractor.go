package keywords

import (
	"strings"
	"unicode/utf8"
)

// 单个实例 线程不安全
type Extractor interface {
	ExtractKeywordsFromHtml(html string) error
	GetResult(filter bool) map[string][]string
	GetContent() string
	GetSubject() string
	Clear()

	Filter(item string) string                      //return filter 后的值
	GetItemsByWeight() (key string, items []string) //按关键字列表序列获取第一个非空结果集
}

type extract struct {
	subject  string                   //关键字主题
	arrKeys  []string                 //关键字列表,有序为权重
	filterFn func(item string) string //GetResult时额外的过滤逻辑

	content string              //解构后的文本
	result  map[string][]string //提取结果
}

var _ Extractor = (*extract)(nil)

type ExtractorOptionFunc func(o *extract)

// arrKeys init
func NewExtractor(keys []string, ops ...ExtractorOptionFunc) Extractor {
	obj := new(extract)
	obj.arrKeys = keys
	obj.result = make(map[string][]string)
	for _, op := range ops {
		op(obj)
	}
	return obj
}

func WithSubject(s string) ExtractorOptionFunc {
	return func(o *extract) {
		o.subject = s
	}
}

var DefaultFilter = func(item string) string {
	//内容过长 大概率是垃圾信息
	if utf8.RuneCountInString(item) > 60 {
		return ""
	}
	//value 包含中文冒号 大概率不是想要的内容
	if strings.ContainsRune(item, ZhColon) {
		return ""
	}
	//
	idx := strings.IndexFunc(item, func(r rune) bool {
		// 此处决定切割关键信息条件
		if r == Sep1 || r == Sep2 || r == Sep3 || r == Sep4 {
			return true
		}
		return false
	})
	if idx > 0 {
		item = item[:idx]
	}
	return item
}

func WithFilter(fn func(string) string) ExtractorOptionFunc {
	return func(o *extract) {
		o.filterFn = fn
	}
}

func (e *extract) GetResult(filter bool) map[string][]string {
	if filter && e.filterFn != nil {
		fr := make(map[string][]string)
		for k, v := range e.result {
			frv := make([]string, 0)
			for _, s := range v {
				r := e.filterFn(s)
				if r != "" {
					frv = append(frv, s)
				}
			}
			if len(frv) > 0 {
				fr[k] = frv
			}
		}
		return fr
	} else {
		return e.result
	}
}

//如果 filterFn 为nil ，返回item本身,无副作用函数
func (e *extract) Filter(item string) string {
	if e.filterFn != nil {
		return e.filterFn(item)
	} else {
		return item
	}
}

func (e *extract) GetContent() string {
	return e.content
}

func (e *extract) GetSubject() string {
	return e.subject
}

func (e *extract) Clear() {
	e.content = ""
	e.result = make(map[string][]string)
}

// 提取关键字
func (e *extract) ExtractKeywordsFromHtml(html string) error {
	content, err := ParseHtml(strings.NewReader(html))
	if err != nil {
		return err
	}
	e.content = content
	for _, key := range e.arrKeys {
		// 返回全部匹配结果
		value := extractKey(content, key)
		if len(value) > 0 {
			e.result[key] = value
		}
	}
	return nil
}

func (e *extract) GetItemsByWeight() (key string, items []string) {
	result := e.GetResult(true)
	for _, key := range e.arrKeys {
		items, ok := result[key]
		if ok == false {
			continue
		}
		return key, items
	}
	return "", nil
}
