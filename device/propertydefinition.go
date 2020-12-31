package device

// PropertyDefinition ...
type PropertyDefinition struct {
	Name  string `bson:"name"`
	Label string `bson:"label"`
	Path  string `bson:"path"`
	Type  string `bson:"type"`
}

// PropertyDefinitionGroup ...
type PropertyDefinitionGroup struct {
	Label               string               `bson:"label"`
	PropertyDefinitions []PropertyDefinition `bson:"propertyDefinitions"`
}
