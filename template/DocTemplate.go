package template

type DocTemplate struct {
	Name  string
	Owner string
	ModelName     string
	ContainerPath string
	Status        string
	Version       string
	FilePath      string
	AttachPath    string
	DocType       string
	Attributes    map[string]string
}
