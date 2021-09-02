package processor

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"path/filepath"
	"pdm.tools/template"
	"strconv"
	"strings"
)

type ImportFile struct {
	filePath  string
	fileType  string
	fileName  string
	modelName string
}

var colMap = map[int]string{
	0: "A",
	1: "B",
	2: "C",
	3: "D",
	4: "E",
	5: "F",
	6: "G",
	7: "H",
	8: "I",
	9: "J",
	10: "K",
	11: "L",
	12: "M",
	13: "N",
	14: "H",
	15: "I",
}

func exportTemplate(exportTemplatePath string, acdDocs []template.DocTemplate){
	f, err := excelize.OpenFile(exportTemplatePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	sheetName := f.GetSheetName(0)

	err = f.InsertRow(sheetName,1)
	if err!=nil{
		fmt.Println(err)
		return
	}



}

func cellAxis(colNum int, rowNum int)(cellAxisStr string){
	colStr := colMap[colNum]
	return colStr+strconv.Itoa(rowNum)
}

func getAcdDocs(importFiles []ImportFile) (acdDocs []template.DocTemplate) {
	for _, importFile := range importFiles {
		acdDoc := template.DocTemplate{
			Name:          importFile.fileName,
			Owner:         "23107",
			ModelName:     importFile.modelName,
			ContainerPath: "/Default",
			Status:        "MODEL_STAGE",
			Version:       "A.1",
			FilePath:      importFile.filePath,
			AttachPath:    "",
			DocType:       importFile.fileType,
		}
		acdDocs = append(acdDocs, acdDoc)
	}
	return acdDocs
}

func fileWalker(path string) (files []string, fileTypes []string, err error) {
	err = filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, filePath)
			fileTypes = append(fileTypes, getSuffix(info.Name()))
		}
		return nil
	})

	return files, fileTypes, err
}

func FileFilter(path string) (importFiles []ImportFile, err error) {
	err = filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			file := ImportFile{
				filePath:  filePath,
				fileType:  getSuffix(info.Name()),
				fileName:  getPrefix(info.Name()),
				modelName: getModelName(filePath),
			}
			importFiles = append(importFiles, file)
		}
		return nil
	})

	return importFiles, err
}

func getSuffix(fileName string) (fileType string) {
	strArray := strings.Split(fileName, ".")
	return strArray[len(strArray)-1]
}

func getPrefix(fileName string) (fileType string) {
	strArray := strings.Split(fileName, ".")
	return strArray[0]
}

func getModelName(filePath string) (modelName string) {
	paths := strings.Split(filePath, "\\")
	return paths[6]
}
