package resources

import (
	"bufio"
	"os"
	"reflect"
	"testing"
)

type ExamContext struct {
}

func Exam(t *testing.T, caseFile string) {
	f, err := os.Open(caseFile)
	if err != nil {
		t.Fatal("open test case failure!")
	}
	defer f.Close()

	var inputs []Case
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		input := convertCase(line)
		if input != nil {
			inputs = append(inputs, *input)
		}
	}

	context := ExamContext{}
	for _, input := range inputs {
		if answer := context.doAnswer(input.Param); !reflect.DeepEqual(answer, input.Answer) {
			t.Fatalf("Fail case is %v; your answer %v ;expected %v", input.Param, answer, input.Answer)
		}
	}
}
