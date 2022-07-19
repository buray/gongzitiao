package excel

import (
	"github.com/xuri/excelize/v2"
	"gongzitiao/myLog"
)

func GetTitleAndData() (title []string, excelData [][]string) {
	var te []string
	var ed [][]string
	f, err := excelize.OpenFile("GZTTZD.xlsx")
	if err != nil {
		myLog.Logger.Printf("打开excel文件失败：%s", err)
		return nil, nil
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			myLog.Logger.Printf("excel文件关闭失败：%s", err)
		}
	}()

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		myLog.Logger.Printf("读取 Sheet1 失败：%s", err)
		return nil, nil
	}

	for s, row := range rows {
		if s == 0 {
			te = row
			continue
		}
		ed = append(ed, row)
	}
	return te, ed
}