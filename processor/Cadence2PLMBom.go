package processor

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"pdm.tools/template"
	"regexp"
	"strconv"
	"strings"
)

const ITEM_NUM_COL = 0
const PART_NUM_COL = 1
const VENDOR_COL = 3
const VENDOR_PART_COL = 4
const QTY_COL = 5
const PLMTEMPLATE_PATH = "D:\\workspace\\go\\src\\pdm.tools\\resource\\BOM\\PLMBOM_template.xlsx"
const BOM_SHEET = "BOM"

func PLMBOMWriter(plmbom template.PLMBOM, path string) error {
	plmbomName := plmbom.AssemblyPartNo + "_" + "PLMBOM.xlsx"
	//plmBOMPath := path+plmbomName
	plmBomLines := plmbom.PLMBomLines
	f, err := excelize.OpenFile(PLMTEMPLATE_PATH)
	if err != nil {
		fmt.Println(err)
		return err

	}
	for rowSize := 0; rowSize < len(plmBomLines); rowSize++ {
		f.InsertRow(BOM_SHEET, 6)
		f.SetCellValue(BOM_SHEET, "A"+strconv.Itoa(6), plmBomLines[len(plmBomLines)-rowSize-1].Item)
		err := f.MergeCell(BOM_SHEET, "I"+strconv.Itoa(6), "J"+strconv.Itoa(6))
		if err != nil {
			return err
		}
		f.SetCellValue(BOM_SHEET, "B"+strconv.Itoa(6), plmBomLines[len(plmBomLines)-rowSize-1].Level)
		f.SetCellValue(BOM_SHEET, "C"+strconv.Itoa(6), plmBomLines[len(plmBomLines)-rowSize-1].CsotPartID)
		f.SetCellValue(BOM_SHEET, "E"+strconv.Itoa(6), plmBomLines[len(plmBomLines)-rowSize-1].Qty)
		f.SetCellValue(BOM_SHEET, "F"+strconv.Itoa(6), plmBomLines[len(plmBomLines)-rowSize-1].Unit)
		f.SetCellValue(BOM_SHEET, "I"+strconv.Itoa(6), plmBomLines[len(plmBomLines)-rowSize-1].UsagePercent)
		f.SetCellValue(BOM_SHEET, "K"+strconv.Itoa(6), plmBomLines[len(plmBomLines)-rowSize-1].ProcedureConsume)
		f.SetCellValue(BOM_SHEET, "L"+strconv.Itoa(6), plmBomLines[len(plmBomLines)-rowSize-1].PartConsume)
		f.SetCellValue(BOM_SHEET, "P"+strconv.Itoa(6), plmBomLines[len(plmBomLines)-rowSize-1].Vendor)
		f.SetCellValue(BOM_SHEET, "Q"+strconv.Itoa(6), plmBomLines[len(plmBomLines)-rowSize-1].VendorPartNo)
	}

	f.SetCellValue(BOM_SHEET, "E4", plmbom.AssemblyPartNo)
	f.SaveAs(plmbomName)
	return nil
}

func CadenceBOM2PLMBOM(cBOM template.CadenceBOM) template.PLMBOM {
	var plmBomLines []template.PLMBOMLine
	cBomLines := cBOM.CBOMLines
	for index, cBomLine := range cBomLines {
		plmBomLine := template.PLMBOMLine{
			Item:             index + 1,
			Level:            1,
			CsotPartID:       cBomLine.CsotPartID,
			Qty:              cBomLine.Qty,
			Unit:             "PCS",
			UsagePercent:     100.0,
			ProcedureConsume: 0.0,
			PartConsume:      0.3,
			Location:         "",
			GeneralFactory:   "General",
			Vendor:           cBomLine.Vendor,
			VendorPartNo:     cBomLine.VendorPartNo,
		}
		plmBomLines = append(plmBomLines, plmBomLine)
	}
	return template.PLMBOM{
		AssemblyPartNo: cBOM.AssemblyPartNo,
		PLMBomLines:    plmBomLines,
	}
}

func GetCadenceBOM(path string) (cBOM template.CadenceBOM) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := f.GetRows("BOM")
	fpcaNo := getFPCANo(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	rowSize := len(rows)

	var startRowNum int
	var lastRowNum int

	for index, row := range rows {
		if row != nil {
			if strings.Compare(strings.ToUpper(row[ITEM_NUM_COL]), "ITEM") == 0 {
				startRowNum = index + 1
			}
		}
	}

	for r := rowSize - 1; r >= 0; r-- {
		firstColcellValue := rows[r][ITEM_NUM_COL]

		if len(firstColcellValue) > 0 {
			_, error := strconv.Atoi(firstColcellValue)
			if error == nil {
				lastRowNum = r + 1
				break
			}
		}
	}
	var cBomLines []template.CadenceBOMLine
	for r := startRowNum; r < lastRowNum; r++ {

		cBomLines = append(cBomLines, template.CadenceBOMLine{
			CsotPartID:   rows[r][PART_NUM_COL],
			Vendor:       rows[r][VENDOR_COL],
			VendorPartNo: rows[r][VENDOR_PART_COL],
			Qty:          rows[r][QTY_COL],
		})
	}

	cBOM = template.CadenceBOM{
		fpcaNo,
		cBomLines,
	}
	return cBOM
}

func getFPCANo(path string) (fpcaNo string) {

	strs := strings.Split(path, string(os.PathSeparator))

	str := strs[len(strs)-1][0:14]

	matched, _ := regexp.MatchString("^34[A-Za-z0-9]{12}$", str)
	if matched {
		fpcaNo = str
	}

	return fpcaNo
}
