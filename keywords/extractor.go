package keywords

import (
	"strings"
)

// 单个实例 线程不安全
type Extractor interface {
	ExtractKeywordsFromHtml(html string) error
	GetResult() map[string]string
	GetContent() string
	GetSubject() string
	Clear()
}

type extract struct {
	Subject string            //关键字主题
	Content string            //解构后的文本
	ArrKeys []string          //关键字列表
	Result  map[string]string //提取结果
}

var _ Extractor = (*extract)(nil)

// ArrKeys init
func NewExtractor(subject string, keys []string) Extractor {
	return &extract{
		Subject: subject,
		ArrKeys: keys,
		Result:  make(map[string]string),
	}
}

func (e *extract) GetResult() map[string]string {
	return e.Result
}

func (e *extract) GetContent() string {
	return e.Content
}

func (e *extract) GetSubject() string {
	return e.Subject
}

func (e *extract) Clear() {
	e.Result = make(map[string]string)
}

// 提取关键字
func (e *extract) ExtractKeywordsFromHtml(html string) error {
	content, err := ParseHtml(strings.NewReader(html))
	if err != nil {
		return err
	}
	e.Content = content
	for _, key := range e.ArrKeys {
		value := extractKey(content, key)
		if value != "" {
			e.Result[key] = value
		}
	}
	return nil
}
