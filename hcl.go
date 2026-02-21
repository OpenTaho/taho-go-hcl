package hcl

// HclType represents the type of an HCL node.
type HclType int

// HclNameAccessor defines an interface for accessing the name of an HCL node.
type HclNameAccessor interface {
	Name() string
}

// HclValueAccessor defines an interface for accessing the value of an HCL node.
type HclValueAccessor interface {
	Value() string
}

// HclNameMutator defines an interface for mutating the name of an HCL node.
type HclNameMutator interface {
	SetName(value string)
}

// HclValueMutator defines an interface for mutating the value of an HCL node.
type HclValueMutator interface {
	SetValue(value string)
}

type HclBodyAccessor interface {
	Body() []HclNode
}

// HclBodyMutator defines an interface for mutating the body of an HCL node.
type HclBodyMutator interface {
	SetBody(value []HclNode)
}

// HclCommentAccessor defines an interface for accessing the comment associated
// with an HCL node.
type HclCommentAccessor interface {
	Comment() string
}

// HclCommentMutator defines an interface for mutating the comment associated
// with an HCL node.
type HclCommentMutator interface {
	SetComment(value string)
}

// HclPairAccessor defines an interface for accessing the pair associated with
// an HCL node.
type HclPairAccessor interface {
	Pair() HclNode
}

// HclPairMutator defines an interface for mutating the pair associated with an
// HCL node.
type HclPairMutator interface {
	SetPair(value HclNode)
}

// HclDir defines an interface for accessing and manipulating HCL directories.
type HclDir interface {
	HclNameAccessor
	HclBodyAccessor
	Files() ([]HclFile, error)
}

// HclBody interface defines the methods for accessing and manipulating the body
// of an HCL node.
type HclBody interface {
	HclBodyAccessor
	HclBodyMutator
}

type HclComment interface {
	HclCommentAccessor
	HclCommentMutator
}

type HclValue interface {
	HclValueAccessor
	HclValueMutator
}

type HclName interface {
	HclNameAccessor
	HclNameMutator
}

type HclPair interface {
	HclPairAccessor
	HclPairMutator
}

// In the context of this parser, an HclNode consists of elements which may in
// turn have subelements. In our model everything associated with a HCL file is
// an HclNode, including the file itself, directories, blocks, pairs, comments,
// and whitespace.
type HclNode interface {
	HclBody
	HclPair
	HclValue
	HclCommentMutator

	AddNode(element HclNode)
	File() HclFile
	IsSimplePair() bool
	Level() int
	Operator() string
	SetDocIndentation(value int)
	SetDocTag(value string)
	SetFileName(value string)
	SetType(value HclType)
	String() string
	Type() HclType
}

// The HclFile interface defines the methods for accessing and manipulating HCL
// files.
type HclFile interface {
	HclNameAccessor
	HclBodyAccessor
	Format() HclFile
	String() string
	SaveAs(fileName string)
}

// The HclParser interface defines the methods for parsing HCL files and
// directories.
type HclParser interface {
	// NewFile creates  a HclFile instances for the associated file name.
	NewFile(name string) HclFile

	// NewDir create a HclDir instances for the associated directory name.
	NewDir(name string) HclDir
}

// String returns the string representation of the HclType.
func (id HclType) String() string {
	return hclElementName[id]
}

// HclType constants represent the different types of HCL nodes.
const (
	HclTypeBlock = iota
	HclTypeComment
	HclTypeDir
	HclTypeDoc
	HclTypeDocWithIndent
	HclTypeOther
	HclTypePair
	HclTypePairGroup
	HclTypeSpace
	HclTypeSpan
	HclTypeString
	HclTypeToken
)

// hclElementName maps HclType constants to their string representations.
var hclElementName = map[HclType]string{
	HclTypeBlock:         "block",
	HclTypeComment:       "comment",
	HclTypeDir:           "dir",
	HclTypeDoc:           "doc",
	HclTypeDocWithIndent: "doc-with-indent",
	HclTypeOther:         "other",
	HclTypePair:          "pair",
	HclTypePairGroup:     "pair-group",
	HclTypeSpace:         "space",
	HclTypeSpan:          "span",
	HclTypeString:        "string",
	HclTypeToken:         "token",
}
