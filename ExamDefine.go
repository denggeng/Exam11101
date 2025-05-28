package resources

import (
	"log"
	"strconv"
	"strings"
)

type Solution interface {
	doAnswer(input Case) Answer
}

// Case 测试用例结构体
type Case struct {
	Param
	Answer
}

// Param 考试题目的输入结构体，请出题人根据实际情况修改此结构体
type Instrument struct {
	p float64
	a int
	b int
}

type Param struct {
	instruments []Instrument
}

// Answer 用户在考试题目返回的结构体，请出题人根据实际情况修改此结构体
type Answer struct {
	result int
}

// 出题人需要实现的函数，用于处理测试用例资源文件返回每一个测试用例
// line为测试用例文件每一行的数据
func convertCase(line string) *Case {
	fields := strings.Fields(line)
	if len(fields) < 1 {
		log.Fatalf("Invalid input: %s", line)
		return nil
	}
	var instruments []Instrument
	for i := 0; i < len(fields)-1; i++ {
		parts := strings.Split(fields[i], ",")
		if len(parts) != 3 {
			log.Fatalf("Invalid instrument format: %s", fields[i])
			return nil
		}
		p, _ := strconv.ParseFloat(parts[0], 64)
		a, _ := strconv.Atoi(parts[1])
		b, _ := strconv.Atoi(parts[2])
		instruments = append(instruments, Instrument{p, a, b})
	}

	ans, _ := strconv.Atoi(fields[len(fields)-1])

	// fmt.Println(instruments, ans)
	return &Case{
		Param:  Param{instruments: instruments},
		Answer: Answer{result: ans},
	}
}
