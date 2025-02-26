package waitfordb

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		numDBs int
	}
	tests := []testCase{
		{1},
		{3},
		{4},
	}

	for _, test := range tests {
		fmt.Printf("---------------------------------\n")
		fmt.Printf("Testing %v Databases...\n\n", test.numDBs)
		dbChan, count := getDBsChannel(test.numDBs)
		waitForDBs(test.numDBs, dbChan)
		for *count != test.numDBs {
			fmt.Println("...")
		}
		fmt.Printf(`
expected length: 0, count: %v
actual length:   %v, count: %v
PASS
---------------------------------
`,
			test.numDBs, len(dbChan), *count)
	}
}
