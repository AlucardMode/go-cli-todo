package structures_tests

import (
	"testing"

	"github.com/AlucardMode/go-cli-todo/structures"
)

type EqualsTooTestCases struct {
	expectedResult bool
	taskA          structures.Task
	taskB          structures.Task
}

func TestEqualsTo(t *testing.T) {
	testCases := []EqualsTooTestCases{
		{
			expectedResult: true,
			taskA:          structures.Task{ID: 0, Description: "test1", State: true},
			taskB:          structures.Task{ID: 0, Description: "test1", State: true},
		},
		// Incorrectly Matching ID
		{
			expectedResult: false,
			taskA:          structures.Task{ID: 0, Description: "test1", State: true},
			taskB:          structures.Task{ID: 1, Description: "test1", State: true},
		},
		// Incorrectly Matching Description
		{
			expectedResult: false,
			taskA:          structures.Task{ID: 0, Description: "test1", State: true},
			taskB:          structures.Task{ID: 0, Description: "test2", State: true},
		},
		// Incorrectly Matching State
		{
			expectedResult: false,
			taskA:          structures.Task{ID: 0, Description: "test1", State: true},
			taskB:          structures.Task{ID: 0, Description: "test1", State: false},
		},
	}

	counter := 0
	for _, testCase := range testCases {
		counter += 1
		if testCase.expectedResult != testCase.taskA.EqualsTo(testCase.taskB) {
			t.Errorf("Test failed for taskA: %+v, taskB: %+v. Expected result: %t", testCase.taskA, testCase.taskB, testCase.expectedResult)

		}
	}

}
