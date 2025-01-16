package postgre

type TableCreationOptions struct {
	IfNotExists   bool
	Temp          bool
	FKConstraints bool
}

type Modeler interface {
	Model(model interface{}) Model
}

type Model interface {
	CreateTable(*TableCreationOptions) error
}
