package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"testing"
	"testing/iotest"
)

func TestOperation(t *testing.T) {
	data := [][]float64{
		{10, 20, 15, 30, 45, 50, 100, 30},
		{5.5, 8, 2.2, 9.75, 8.45, 3, 2.5, 10.25, 4.75, 6.1, 7.67, 12.287, 5.47},
		{-10, -20},
		{102, 37, 44, 57, 67, 129},
	}

	//TestCases for operation test
	testCases := []struct {
		name string
		op   statsFunc
		exp  []float64
	}{
		{"Sum", sum, []float64{300, 85.927, -30, 436}},
		{"Avg", avg, []float64{37.5, 6.609769230769231, -15, 72.666666666666666}},
	}

	//Operations Tests Execution
	for _, tc := range testCases {
		for k, exp := range tc.exp {
			name := fmt.Sprintf("%sData%d", tc.name, k)
			t.Run(name, func(t *testing.T) {
				result := tc.op(data[k])
				if result != exp {
					t.Errorf("Expected %g, got %g instead", exp, result)
				}
			})
		}
	}
}

func TestCSV2Float(t *testing.T) {
	csvData := `
	IP Address,Requests,Response Time
	192.168.0.199,2056,236
	192.168.0.88,899,220
	192.168.0.199,3054,226
	192.168.0.100,4133,218
	192.168.0.199,950,238`

	//TestCases for csv2float test
	testCases := []struct {
		name     string
		col      int
		exp      []float64
		expErr	 error
		r        io.Reader
	}{
		{name: "Column2", col: 2,
			exp:    []float64{2056, 899, 3054, 4133, 950},
			expErr: nil,
			r:      bytes.NewBufferString(csvData),
		},
		{name: "Column3", col: 3,
			exp:    []float64{236, 220, 226, 218, 238},
			expErr: nil,
			r:      bytes.NewBufferString(csvData),
		},
		{name: "FailRead", col: 1,
			exp:    nil,
			expErr: iotest.ErrTimeout,
			r:      iotest.TimeoutReader(bytes.NewReader([]byte{0})),
		},
		{name: "FailedNotNumber", col: 1,
			exp:    nil,
			expErr: ErrNotNumber,
			r:      bytes.NewBufferString(csvData),
		},
		{name: "FailedInvalidColumn", col: 4,
			exp:    nil,
			expErr: ErrInvalidColumn,
			r:      bytes.NewBufferString(csvData),
		},
	}

	//csv2float test execution
	for _, tc := range testCases{
		t.Run(tc.name, func(t *testing.T) {
			res, err := csv2float(tc.r, tc.col)

			//check for errors if expErr is not nil
			if tc.expErr != nil {
				if err == nil {
					t.Errorf("Expected error, got nil instead")
				}
				if !errors.Is(err, tc.expErr){
					t.Errorf("expected error %q, got %q instead", tc.expErr, err)
				}
				return
			}

			//check results if errors are not expected
			if err != nil {
				t.Errorf("Unexpected error: %q", err)
			}

			for i, exp := range tc.exp {
				if res[i] != exp {
					t.Errorf("Expected %g, got %g instead", exp, res[i])
				}
			}			
		})
	}
}
