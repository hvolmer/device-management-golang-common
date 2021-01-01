package compliance

import (
	"log"

	"github.com/hvolmer/device-management-golang-common/device"
	"github.com/hvolmer/device-management-golang-common/devicechange"
	"github.com/xeipuuv/gojsonschema"
)

// GetComplianceForChange ...
func (dev device.Device) GetComplianceForChange(ch devicechange.DeviceChange) string {
	for _, ci := range dev.ComplianceItems {
		if ci.Path == ch.Path {
			schemaLoader := gojsonschema.NewGoLoader(ci.Schema)
			documentLoader := gojsonschema.NewGoLoader(ch.Value)
			result, err := gojsonschema.Validate(schemaLoader, documentLoader)
			if err != nil {
				log.Println("Error validating compliance. [value]", ch.Value, "[schema]", ci.Schema)
			}
			log.Println("COMPLIANCE RESULT:", result)
		}
	}
	return ""
}
