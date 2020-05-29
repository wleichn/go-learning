package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const chinese = "Chinese"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const chineseHelloPrefix = "你好， "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case chinese:
		prefix = chineseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}

/**
总结：
1、函数命名返回值：会在函数中创建一个名为prefix的变量，会被分配零值，字符串初始值为 ""；只需要调用return即可返回所设置的值
2、函数名称以小写字母开头，在Go中公共函数以大写字符开头，私有函数以小写字母开头
3、TDD：测试驱动开发
*/
