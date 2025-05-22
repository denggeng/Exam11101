package resources

import (
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
type Param struct {
	arr []int
	n   int
}

// Answer 用户在考试题目返回的结构体，请出题人根据实际情况修改此结构体
type Answer struct {
	rs1 int
}

// 出题人需要实现的函数，用于处理测试用例资源文件返回每一个测试用例
// line为测试用例文件每一行的数据
func convertCase(line string) *Case {

	sArr := strings.Split(line, " ")
	if len(sArr) == 3 {
		arr, _ := stringToArray(sArr[0])
		n, _ := strconv.Atoi(sArr[1])
		rs1, _ := strconv.Atoi(sArr[2])
		return &Case{Param{arr: arr, n: n}, Answer{rs1}}
	}
	return nil
}
func stringToArray(s string) ([]int, error) {
	parts := strings.Split(s, ",")
	arr := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			return nil, err
		}
		arr[i] = num
	}
	return arr, nil
}
