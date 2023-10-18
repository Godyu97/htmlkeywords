package keywords

import (
	"strings"
)

// 单个实例 线程不安全
type Extractor interface {
	ExtractKeywordsFromHtml(html string) error
	GetResult(filter bool) map[string]string
	GetContent() string
	GetSubject() string
	Clear()

	Filter(item string) string //return filter 后的值
}

type extract struct {
	subject  string                   //关键字主题
	arrKeys  []string                 //关键字列表
	filterFn func(item string) string //GetResult时额外的过滤逻辑

	content string            //解构后的文本
	result  map[string]string //提取结果
}

var _ Extractor = (*extract)(nil)

type ExtractorOptionFunc func(o *extract)

// arrKeys init
func NewExtractor(keys []string, ops ...ExtractorOptionFunc) Extractor {
	obj := new(extract)
	obj.arrKeys = keys
	obj.result = make(map[string]string)
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

func WithFilter(fn func(string) string) ExtractorOptionFunc {
	return func(o *extract) {
		o.filterFn = fn
	}
}

func (e *extract) GetResult(filter bool) map[string]string {
	if filter && e.filterFn != nil {
		fr := make(map[string]string, len(e.result))
		for k, v := range e.result {
			frv := e.filterFn(v)
			if frv != "" {
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
	e.result = make(map[string]string)
}

// 提取关键字
func (e *extract) ExtractKeywordsFromHtml(html string) error {
	content, err := ParseHtml(strings.NewReader(html))
	if err != nil {
		return err
	}
	e.content = content
	for _, key := range e.arrKeys {
		value := extractKey(content, key)
		if value != "" {
			e.result[key] = value
		}
	}
	return nil
}
