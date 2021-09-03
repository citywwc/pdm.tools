package template

type PLMBOM struct {
	AssemblyPartNo string
	PLMBomLines    []PLMBOMLine
}

type PLMBOMLine struct {
	Item             int
	Level            int
	CsotPartID       string
	Qty              string
	Unit             string
	UsagePercent     float32
	ProcedureConsume float32
	PartConsume      float32
	Location         string
	GeneralFactory   string
	Vendor           string
	VendorPartNo     string
}

func PLMBOMInit(item int, csotPartID string, qty string, location string, vendor string, vendorPartNo string) *PLMBOMLine {
	return &PLMBOMLine{
		Item:             item,
		Level:            1,
		CsotPartID:       csotPartID,
		Qty:              qty,
		Unit:             "PCS",
		UsagePercent:     100.0,
		ProcedureConsume: 0.0,
		PartConsume:      0.3,
		Location:         location,
		GeneralFactory:   "General",
		Vendor:           vendor,
		VendorPartNo:     vendorPartNo,
	}
}
