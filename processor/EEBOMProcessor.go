package processor

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const HUAXINPARTNOCOL = 2
const MANUFACTURE = 3
const MANUFACTUREPARTNOCOL = 4

var PartsMapping = make(map[string]interface{})

func BOMFileReader(path string) (bomFilePaths []string, err error) {
	err = filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			bomFilePaths = append(bomFilePaths, filePath)
		}
		return nil
	})
	return bomFilePaths, err
}

func BOMWriter(path string, partMPartMapping map[string]interface{}) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := f.GetRows("BOM")
	if err != nil {
		fmt.Println(err)
		return
	}

	for rowIndex := 5; rowIndex < len(rows); rowIndex++ {
		partNumber := rows[rowIndex][2]
		interfaceType := partMPartMapping[partNumber]
		manufactureMapping := interfaceType.(map[string]string)
		mParts := ""
		manufactures := ""
		for mPart, manufacture := range manufactureMapping {
			mParts = mParts + mPart + "\n"
			manufactures = manufactures + manufacture + "\n"
		}
		f.SetCellDefault("BOM", "P"+strconv.Itoa(rowIndex), manufactures)
		f.SetCellDefault("BOM", "Q"+strconv.Itoa(rowIndex), mParts)
	}
	f.Save()
}
func ManufactureMappingReader(path string) (manufactureMapping map[string]string) {
	manufactureMapping = make(map[string]string)
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rowIndex := 1; rowIndex < len(rows); rowIndex++ {
		manufactureMapping[rows[rowIndex][MANUFACTUREPARTNOCOL]] = rows[rowIndex][MANUFACTURE]
	}
	return manufactureMapping
}

func HxPartMPartMappingReader(path string, manufactureMapping map[string]string) (partMPartMapping map[string]interface{}) {
	partMPartMapping = make(map[string]interface{})
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	for rowIndex1 := 1; rowIndex1 < len(rows); rowIndex1++ {
		partNumber := rows[rowIndex1][HUAXINPARTNOCOL]
		mpNumber := rows[rowIndex1][MANUFACTUREPARTNOCOL]
		mPartSupplierMapping := make(map[string]string)
		if partMPartMapping[partNumber] != nil {
			interfaceType := partMPartMapping[partNumber]
			mPartSupplierMapping = interfaceType.(map[string]string)
		}

		for rowIndex2 := 1; rowIndex2 < len(rows); rowIndex2++ {
			mPartSupplierMapping[mpNumber] = manufactureMapping[mpNumber]
			partMPartMapping[partNumber] = mPartSupplierMapping
			nextPartNumber := rows[rowIndex2][HUAXINPARTNOCOL]
			if strings.Compare(nextPartNumber, partNumber) == 0 {
				mPartSupplierMapping[mpNumber] = manufactureMapping[mpNumber]
			}
		}

	}

	return partMPartMapping
}
