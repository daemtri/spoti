package apispot

type Specific struct {
}

type Struct struct {
	PackageName  string    `json:"packageName"`
	Filename     string    `json:"filename"`
	DocLines     []string  `json:"docLines,omitempty"`
	Name         string    `json:"name"`
	Fields       []Field   `json:"fields,omitempty"`
	Operations   []*Method `json:"operations,omitempty"`
	CommentLines []string  `json:"commentLines,omitempty"`
}

type Field struct {
	PackageName  string   `json:"packageName,omitempty"`
	DocLines     []string `json:"docLines,omitempty"`
	Name         string   `json:"name,omitempty"`
	TypeName     string   `json:"typeName,omitempty"`
	Tag          string   `json:"tag,omitempty"`
	CommentLines []string `json:"commentLines,omitempty"`
}

type Interface struct {
	Filename     string   `json:"filename"`
	DocLines     []string `json:"docLines,omitempty"`
	Name         string   `json:"name"`
	Methods      []Method `json:"methods,omitempty"`
	CommentLines []string `json:"commentLines,omitempty"`
}

type Method struct {
	PackageName   string   `json:"packageName,omitempty"`
	Filename      string   `json:"filename,omitempty"`
	DocLines      []string `json:"docLines,omitempty"`
	RelatedStruct *Field   `json:"relatedStruct,omitempty"` // optional
	Name          string   `json:"name"`
	InputArgs     []Field  `json:"inputArgs,omitempty"`
	OutputArgs    []Field  `json:"outputArgs,omitempty"`
	CommentLines  []string `json:"commentLines,omitempty"`
}

type Errors struct {
}

type Enum struct {
}
