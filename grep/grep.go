package grep

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sync"
)

var lines = make(chan string) // 创建一个通道来传递匹配的行

func Grep(pattern string, filename string, concurrentNum int) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var wg sync.WaitGroup

	// 启动多个goroutine来处理每一行
	for i := 0; i < concurrentNum; i++ { // 这里可以根据需求调整并发数
		wg.Add(1)
		go func(threadNum int) {
			defer wg.Done()
			// TODO: 为什么bufio.NewScanner(file).scanner.Scan()能支持并发扫描文件且不重复？
			for scanner.Scan() { // 逐行扫描文件
				line := scanner.Text()
				//fmt.Printf("Thread %d: %s\n", threadNum, line)
				putMatchString2Chan(pattern, line)
				//time.Sleep(3 * time.Second)
			}
		}(i + 1)
	}

	// 启动一个goroutine来关闭通道并等待所有线程完成
	go func() {
		wg.Wait()
		//fmt.Println("close(lines)")
		close(lines)
	}()

	// 从通道中读取并打印匹配的行
	for line := range lines {
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func putMatchString2Chan(pattern string, line string) {
	if matched, _ := regexp.MatchString(pattern, line); matched {
		//fmt.Println(line)
		lines <- line // 将匹配的行发送到通道
	}
}
