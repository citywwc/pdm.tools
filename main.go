package main

import (
	"pdm.tools/processor"
)

type testFile struct {
	FilePath  string
	FileType  string
	FileName  string
	ModelName string
}

func main() {

	path := "D:\\workspace\\go\\src\\pdm.tools\\resource\\BOM\\mapping.xlsx"
	BOMPath := "D:\\workspace\\go\\src\\pdm.tools\\resource\\BOM\\34421100248041.xls"
	manufactureMapping := processor.ManufactureMappingReader(path)
	partMPartMapping := processor.HxPartMPartMappingReader(path, manufactureMapping)

	//for _, value := range partMPartMapping {
	//	manufactureMapping := value.(map[string]string)
	//	fmt.Println(manufactureMapping)
	//	mParts := ""
	//	manufactures := ""
	//	for mPart, manufacture := range manufactureMapping {
	//		mParts = mParts + mPart + "\n"
	//		manufactures = manufactures + manufacture + "\n"
	//	}
	//	fmt.Println("=========")
	//	fmt.Print(mParts)
	//	fmt.Println("---------")
	//	fmt.Print(manufactures)
	//	fmt.Println("=========")
	//	break
	//}
	processor.BOMWriter(BOMPath,partMPartMapping)

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
