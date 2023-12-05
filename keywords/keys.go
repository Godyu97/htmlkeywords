package keywords

import (
	"github.com/Godyu97/vege9/vege"
	"regexp"
	"strings"
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

func regKey(content string, key []rune) (after string, found bool) {
	pattern := ""
	for i := 0; i < len(key); i++ {
		//	处理小括号
		if key[i] == '(' || key[i] == ')' {
			pattern += `\`
		}
		if i < len(key)-1 {
			pattern += string(key[i]) + KeyP
		} else {
			pattern += string(key[i])
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

func extractKey(content, key string) []string {
	result := make([]string, 0)
	//  实现字符串匹配截取关键内容 关键字可能有空格 正则取index
	after, found := regKey(content, []rune(key))
	//contain key
	if found == true {
		text := []rune(after)
		//text 不为空
		if len(text) > 0 {
			// 1. 处理冒号之后的内容 如果key后面直接是文本，说明不是需要的信息 继续向下匹配
			if text[0] != ZhColon && text[0] != EnColon && text[0] != Enter && text[0] != Space && text[0] != Table && text[0] != ENSP && text[0] != EMSP && text[0] != NBSP {
				return append(result, extractKey(after, key)...)
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
				return result
			}
			// 3. 有效内容，截取到终止条件
			s = s[idxText:]
			idx := strings.IndexFunc(s, func(r rune) bool {
				// 此处决定切割关键信息条件
				if r == Enter {
					return true
				}
				return false
			})
			if idx > 0 {
				s = s[:idx]
			}
			//去除不成对的括号
			s = vege.RemoveInvalidParentheses(s, [2]rune{'(', ')'})
			s = vege.RemoveInvalidParentheses(s, [2]rune{'（', '）'})
			//去除末端空格
			s = strings.TrimSpace(s)
			if s != "" {
				result = append(result, s)
			}
			return append(result, extractKey(after, key)...)
		}
	}
	//no contain
	return result
}
