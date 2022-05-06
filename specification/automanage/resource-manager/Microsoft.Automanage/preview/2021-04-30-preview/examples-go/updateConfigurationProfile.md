Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fautomanage%2Farmautomanage%2Fv0.4.0/sdk/resourcemanager/automanage/armautomanage/README.md) on how to add the SDK to your project and authenticate.

```go
package armautomanage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automanage/armautomanage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automanage/resource-manager/Microsoft.Automanage/preview/2021-04-30-preview/examples/updateConfigurationProfile.json
func ExampleConfigurationProfilesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armautomanage.NewConfigurationProfilesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<configuration-profile-name>",
		"<resource-group-name>",
		armautomanage.ConfigurationProfileUpdate{
			Tags: map[string]*string{
				"Organization": to.Ptr("Administration"),
			},
			Properties: &armautomanage.ConfigurationProfileProperties{
				Configuration: map[string]interface{}{
					"Antimalware/Enable":                false,
					"AzureSecurityCenter/Enable":        true,
					"Backup/Enable":                     false,
					"BootDiagnostics/Enable":            true,
					"ChangeTrackingAndInventory/Enable": true,
					"GuestConfiguration/Enable":         true,
					"LogAnalytics/Enable":               true,
					"UpdateManagement/Enable":           true,
					"VMInsights/Enable":                 true,
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
