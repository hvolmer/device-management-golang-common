package compliance

import (
	"log"

	"github.com/hvolmer/device-management-golang-common/device"
	"github.com/hvolmer/device-management-golang-common/devicechange"
	"github.com/xeipuuv/gojsonschema"
)

// GetComplianceForChange ...
func GetComplianceForChange(ch devicechange.DeviceChange, compItems *[]device.ComplianceItem) string {
	for _, ci := range *compItems {
		if ci.Path == ch.Path {
			schemaLoader := gojsonschema.NewGoLoader(ci.Schema)
			documentLoader := gojsonschema.NewGoLoader(ch.Value)
			result, err := gojsonschema.Validate(schemaLoader, documentLoader)
			if err != nil {
				log.Println("Error validating compliance. [value]", ch.Value, "[schema]", ci.Schema)
			}
			if result.Valid() {
				log.Println("COMPLIANCE ES BUENO")
				return "false"
			} else {
				log.Println("COMPLIANCE RESULT:", result.Errors())
				return "true"
			}
		}
	}
	return ""
}
