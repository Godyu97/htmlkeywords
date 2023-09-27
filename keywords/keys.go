package keywords

import (
	"github.com/Godyu97/vege9/vegeTools"
	"regexp"
	"strings"
	"unicode/utf8"
)

const (
	ZhColon = '：'
	EnColon = ':'
	//结束提取字符
	ENSP  = '\u2002'
	EMSP  = '\u2003'
	NBSP  = '\xa0'
	Space = '\x20'
	Enter = '\n'
	Table = '\t'
	Sep1  = '，'
	Sep2  = '。'
	Sep3  = '；'
	Sep4  = '】'
)

// 关键字之间处理空格的正则段 32 or 160
const KeyP = `(\s|\xa0)*`

func regKey(content string, key string) (after string, found bool) {
	pattern := ""
	for i, r := range key {
		pattern += string(r)
		if i < utf8.RuneCountInString(key)-1 {
			pattern += KeyP
		}
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", false
	}
	match := re.FindStringIndex(content)
	if match != nil {
		return content[match[1]:], true
	}
	return "", false
}

func extractKey(content, key string) string {
	//  实现字符串匹配截取关键内容 关键字可能有空格 正则取index
	after, found := regKey(content, key)
	//contain key
	if found == true {
		text := []rune(after)
		//text 不为空
		if len(text) > 0 {
			// 1. 处理冒号之后的内容 如果key后面直接是文本，说明不是需要的信息 继续向下匹配
			if text[0] != ZhColon && text[0] != EnColon && text[0] != Enter && text[0] != Space && text[0] != Table && text[0] != ENSP && text[0] != EMSP && text[0] != NBSP {
				return extractKey(after, key)
			}
			text = text[1:]
			s := string(text)
			// 2. 去掉开头冒号后的空格和回车
			idxText := strings.IndexFunc(s, func(r rune) bool {
				if r != Space && r != Enter && r != Table && r != ENSP && r != EMSP && r != NBSP {
					return true
				}
				return false
			})
			if idxText < 0 {
				//全是空的，无关键内容
				return ""
			}
			// 3. 有效内容，截取到终止条件
			s = s[idxText:]
			idx := strings.IndexFunc(s, func(r rune) bool {
				//todo 此处决定切割关键信息条件
				if r == Enter ||
					r == Sep1 || r == Sep2 || r == Sep3 || r == Sep4 {
					return true
				}
				return false
			})
			if idx > 0 {
				s = s[:idx]
			}
			//内容过长 大概率是垃圾信息 去after 找找有没有有效信息
			if utf8.RuneCountInString(s) > 60 {
				return extractKey(after, key)
			}
			//value 包含中文冒号 大概率不是想要的内容
			if strings.ContainsRune(s, ZhColon) {
				return extractKey(after, key)
			}
			//去除不成对的括号
			s = vegeTools.RemoveInvalidParentheses(s, [2]rune{'(', ')'})
			s = vegeTools.RemoveInvalidParentheses(s, [2]rune{'（', '）'})
			//去除末端空格
			s = strings.TrimSpace(s)
			return s
		}
	}
	//no contain
	return ""
}
