package excel

import (
	"fmt"
	"testing"
)

func TestGetTitleAndData(t *testing.T) {
	title, excelData := GetTitleAndData()
	fmt.Println(title, excelData)
}