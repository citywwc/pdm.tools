package main

import (
	"os"
	"path/filepath"
	"pdm.tools/processor"
)

type testFile struct {
	FilePath  string
	FileType  string
	FileName  string
	ModelName string
}

func fileWalker(path string) (files []string, err error) {
	err = filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, filePath)

		}
		return nil
	})
	return files, err
}

func main() {

	//path := "D:\\workspace\\go\\src\\pdm.tools\\resource\\BOM\\mapping.xlsx"
	cBOMFolderPath := "D:\\workspace\\go\\src\\pdm.tools\\resource\\BOM\\"
	//cBOMPath := "E:\\CSOT\\电子料梳理\\34298X00370211_PCBA_BOM_V01_20210111.xlsx"
	//manufactureMapping := processor.ManufactureMappingReader(path)
	//partMPartMapping := processor.HxPartMPartMappingReader(path, manufactureMapping)

	//processor.BOMWriter(BOMPath,partMPartMapping)
	cBOMFiles, _ := fileWalker(cBOMFolderPath)
	for _, cBOMPath := range cBOMFiles {
		cBOM := processor.GetCadenceBOM(cBOMPath)
		plmBOM := processor.CadenceBOM2PLMBOM(cBOM)
		processor.PLMBOMWriter(plmBOM, "")
	}

	//mapTest := map[string]interface{}{
	//	"aa": "aa",
	//	"bb": "bb",
	//}
	//
	//if mapTest["cc"] == nil {
	//	fmt.Println("!!!")
	//}

	//acdDoc := testFile{
	//	FilePath: "GXMD",
	//	FileType: "Dwg",
	//	FileName: "testfile",
	//}
	//
	//
	//
	//t := reflect.TypeOf(acdDoc)
	//v := reflect.ValueOf(acdDoc)
	//
	//for k := 0; k < t.NumField(); k++{
	//	fmt.Printf("%s -- %v \n", t.Field(k).Name, v.Field(k).Interface())
	//}
	//
	//ManufacturePartMapping := map[string]string{
	//	"CL03A105MP3NSNC": "三星",
	//	"LDK063ABJ105MP-F": "太诱",
	//}
	//
	////partMapping := map[string]interface{}{
	////	"34050001U1X2M1": ManufacturePartMapping,
	////}
	//
	//for key,value := range ManufacturePartMapping{
	//	fmt.Println(key + " --- "+value)
	//}
}
