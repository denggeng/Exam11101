package resources

import (
	"bufio"
	"os"
	"reflect"
	"testing"
	// "fmt"
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
	for i, input := range inputs {
		if answer := context.doAnswer(input.Param); !reflect.DeepEqual(answer, input.Answer) {
			t.Fatalf("Test case %d failed: input = %v; actual = %v; expected = %v", i, input, answer, input.Answer)
		}
	}
}
