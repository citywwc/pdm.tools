package template

type CadenceBOM struct {
	AssemblyPartNo string
	CBOMLines      []CadenceBOMLine
}

type CadenceBOMLine struct {
	CsotPartID   string
	Vendor       string
	VendorPartNo string
	Qty          string
	Location     string
}
