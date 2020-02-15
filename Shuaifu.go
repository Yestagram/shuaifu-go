package shuaifu

import "fmt"

var firstLetter = map[rune]string{
	'b': "b", 'c': "c", 'd': "d", 'f': "f", 'g': "g", 'h': "h", 'i': "ch", 'j': "j", 'k': "k", 'l': "l",
	'm': "m", 'n': "n", 'p': "p", 'q': "q", 'r': "r", 's': "s", 't': "t",
	'u': "sh", 'v': "zh", 'w': "w", 'x': "x", 'y': "y", 'z': "z"}
var secondLetter = map[rune]string{
	'a': "a", 'b': "in", 'c': "ao", 'd': "ai",
	'e': "e", 'f': "en", 'g': "eng", 'h': "ang",
	'i': "i", 'j': "an", 'm': "ian", 'n': "iao",
	'p': "ie", 'q': "iu", 'r': "uan",
	'u': "u", 'w': "ei", 'y': "un", 'z': "ou"}
var secondLetterMul = map[rune][2]string{
	'k': {"ing", "uai"}, 'l': {"iang", "uang"}, 'o': {"o", "uo"}, 's': {"iong", "ong"},
	't': {"ue", "ve"}, 'v': {"ui", "v"}, 'x': {"ia", "ua"},
}
var otherLetter = map[string]string{
	"aa": "a", "ah": "ang", "ai": "ai", "an": "an", "ao": "ao", "ee": "e",
	"eg": "eng", "ei": "ei", "en": "en", "er": "er", "oo": "o", "ou": "ou"}
var validWords = []string{"jia", "qia", "xia", "gua", "kua", "hua", "zhua", "chua", "shua", "guai", "kuai", "huai",
	"zhuai", "chuai", "shuai", "bing", "ping", "ming", "ding", "ting", "ning", "ling", "jing", "qing",
	"xing", "ying", "guang", "kuang", "huang", "zhuang", "chuang", "shuang", "diang", "niang", "liang",
	"jiang", "qiang", "xiang", "bo", "po", "mo", "fo", "lo", "wo", "duo", "tuo", "nuo", "luo", "guo", "kuo",
	"huo", "zhuo", "chuo", "shuo", "ruo", "zuo", "cuo", "suo", "jiong", "qiong", "xiong", "dong", "tong",
	"nong", "long", "gong", "kong", "hong", "zhong", "chong", "rong", "zong", "cong", "song", "yong",
	"jue", "que", "xue", "yue", "nve", "lve", "dui", "tui", "gui", "kui", "hui", "zhui", "chui", "shui", "rui",
	"zui", "cui", "sui", "nv", "lv"}

func hasStr(arr []string, needle string) bool {
	for _, e := range arr {
		if e == needle {
			return true
		}
	}
	return false
}

func Believe() {
	fmt.Println("我信仰帅副!")
}

func Version() string {
	return "0.1.0"
}

func Translate(text string, opt ...bool) string {
	var result = ""
	var tmp = []rune(text)
	var length = len(tmp)
	var addBlank = len(opt) == 0 || len(opt) > 0 && opt[0]
	for i := 0; i < length; i++ {
		if i == length-1 {
			result += string(tmp[i])
			break
		}
		var t = string(tmp[i]) + string(tmp[i+1])
		var blank = ""
		if tmp[i] != ' ' && addBlank {
			blank = " "
		}
		if ol, ok := otherLetter[t]; ok {
			result += ol + blank
			i++
		} else {
			var t1, ok = firstLetter[tmp[i]]
			var t2, ok2 = secondLetter[tmp[i+1]]
			var t22, ok22 = secondLetterMul[tmp[i+1]]
			if ok && (ok2 || ok22) {
				if ok22 {
					var w1 = firstLetter[tmp[i]] + t22[0]
					var w2 = firstLetter[tmp[i]] + t22[1]
					if hasStr(validWords, w1) {
						result += w1 + blank
					} else if hasStr(validWords, w2) {
						result += w2 + blank
					}
				} else {
					result += t1 + t2 + blank
				}
				i++
			} else {
				result += string(tmp[i])
			}
		}
	}
	return result
}
