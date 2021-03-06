package armiotsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotsecurity/armiotsecurity"
)

// x-ms-original-file: specification/iotsecurity/resource-manager/Microsoft.IoTSecurity/preview/2021-07-01-preview/examples/Alerts/PatchAlert.json
func ExampleIotAlertsClient_Patch() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiotsecurity.NewIotAlertsClient("<subscription-id>",
		"<iot-defender-location>", cred, nil)
	res, err := client.Patch(ctx,
		"<device-group-name>",
		"<alert-id>",
		armiotsecurity.AlertPatchPropertiesModel{
			Properties: &armiotsecurity.AlertPatchPropertiesModelProperties{
				Severity: armiotsecurity.AlertSeverity("Medium").ToPtr(),
				Status:   armiotsecurity.AlertStatus("InProgress").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IotAlertsClientPatchResult)
}
