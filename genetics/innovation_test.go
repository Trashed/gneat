package genetics_test

import (
	"fmt"
	"testing"

	"github.com/Trashed/gneat/genetics"
	"github.com/Trashed/gneat/internal/assertions"
)

func TestAddConnection_String_Success(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		genes             []string
		expectedInnovNums []uint
	}{
		{
			name:              "unique genes",
			genes:             []string{"1->4", "2->4", "4->3"},
			expectedInnovNums: []uint{1, 2, 3},
		},
		{
			name:              "similar genes",
			genes:             []string{"1->4", "2->4", "2->4", "4->5", "1->4"},
			expectedInnovNums: []uint{1, 2, 2, 3, 1},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestAddConnection_String_%s_%d", test.name, i), func(t *testing.T) {
			innovationRecord := genetics.NewInnovationRecord[string]()

			for j, g := range test.genes {
				actualInnovNum, _ := innovationRecord.Store(g)

				assertions.Equals(t, test.expectedInnovNums[j], actualInnovNum)
			}
		})
	}
}

func TestAddConnection_String_Failure(t *testing.T) {
	t.Parallel()

	innovationRecord := genetics.NewInnovationRecord[string]()
	const gene = ""
	const expectedInnovNum uint = 0
	var expectedErr error = genetics.ErrInvalidGeneValue

	actualInnovNum, err := innovationRecord.Store(gene)
	assertions.Assert(t, err != nil, "expected an error", expectedErr)
	assertions.Equals(t, expectedInnovNum, actualInnovNum)
}
