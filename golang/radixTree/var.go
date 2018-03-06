package radixTree

// charset 为字符集合,该字符集可以根据实际情况来增加或者减少
var charset []byte
var charsetLen int

func init() {
	charset = []byte{
		// 小写字母
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'o', 'p',
		'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		// 数字
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		// 特殊字符
		'/',
	}
	charsetLen = len(charset)
}

func getCharsetIndex(charset byte) (ret int) {
	switch charset {
	case 'a':
		ret = 0
	case 'b':
		ret = 1
	case 'c':
		ret = 2
	case 'd':
		ret = 3
	case 'e':
		ret = 4
	case 'f':
		ret = 5
	case 'g':
		ret = 6
	case 'h':
		ret = 7
	case 'i':
		ret = 8
	case 'j':
		ret = 9
	case 'k':
		ret = 10
	case 'l':
		ret = 11
	case 'm':
		ret = 12
	case 'n':
		ret = 13
	case 'o':
		ret = 14
	case 'p':
		ret = 15
	case 'q':
		ret = 16
	case 'r':
		ret = 17
	case 's':
		ret = 18
	case 't':
		ret = 19
	case 'u':
		ret = 20
	case 'v':
		ret = 21
	case 'w':
		ret = 22
	case 'x':
		ret = 23
	case 'y':
		ret = 24
	case 'z':
		ret = 25
	case '0':
		ret = 26
	case '1':
		ret = 27
	case '2':
		ret = 28
	case '3':
		ret = 29
	case '4':
		ret = 30
	case '5':
		ret = 31
	case '6':
		ret = 32
	case '7':
		ret = 33
	case '8':
		ret = 34
	case '9':
		ret = 35
	case '/':
		ret = 36
	default:
		ret = -1
	}
	return ret
}
