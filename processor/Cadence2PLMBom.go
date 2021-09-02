package processor

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"pdm.tools/template"
	"reflect"
	"strconv"
	"strings"
)

const ITEM_NUM_COL = 0
const PART_NUM_COL = 1
const VENDOR_COL = 3
const VENDOR_PART_COL = 4
const QTY_COL = 5

func GetCadenceBOMLine(path string) (cBomLines []template.CadenceBOMLine) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := f.Rows("BOM")

	if err != nil {
		fmt.Println(err)
		return
	}
	vRows := reflect.ValueOf(rows)
	rowSize := vRows.FieldByName("totalRow").Int()

	var startRowNum int
	var lastRowNum int

	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			fmt.Println(err)
		}
		for _, colCell := range row {
			if strings.Compare(strings.ToUpper(colCell), "ITEM") == 0 {
				startRowNum = int(vRows.FieldByName("curRow").Int())
			}
		}
	}
	//TODO Get last row num
	//for r := rowSize - 1; r >= 0; r-- {
	//	firstColcellValue := rows[r][ITEM_NUM_COL]
	//
	//	if len(firstColcellValue) > 0 {
	//		_, error := strconv.Atoi(firstColcellValue)
	//		if error == nil {
	//			lastRowNum = r
	//			break
	//		}
	//	}
	//}
	//TODO set bom lines collections

	//	for r := startRowNum; r < lastRowNum; r++ {
	//
	//		cBomLines = append(cBomLines, template.CadenceBOMLine{
	//			CsotPartID:   rows[r][PART_NUM_COL],
	//			Vendor:       rows[r][VENDOR_COL],
	//			VendorPartNo: rows[r][VENDOR_PART_COL],
	//			Qty:          rows[r][QTY_COL],
	//		})
	//	}
	return cBomLines
}
