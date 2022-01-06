package interfaces

type ObjectType string

type IMetadata interface {

	// Set
	SetNamespace(string)
	SetName(string)
	SetKind(string)
	SetWorkload(map[string]interface{}) // DEPRECATED
	SetObject(map[string]interface{})
	// TODO - AetApiVersion

	// Get
	GetNamespace() string
	GetName() string
	GetKind() string
	GetApiVersion() string
	GetWorkload() map[string]interface{} // DEPRECATED
	GetObject() map[string]interface{}
	GetID() string // Get object unique ID

	GetObjectType() ObjectType // Get struct type

}
