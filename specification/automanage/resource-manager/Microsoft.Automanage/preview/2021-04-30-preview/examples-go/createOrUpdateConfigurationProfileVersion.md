```go
package armautomanage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automanage/armautomanage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automanage/resource-manager/Microsoft.Automanage/preview/2021-04-30-preview/examples/createOrUpdateConfigurationProfileVersion.json
func ExampleConfigurationProfilesVersionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armautomanage.NewConfigurationProfilesVersionsClient("mySubscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"customConfigurationProfile",
		"version1",
		"myResourceGroupName",
		armautomanage.ConfigurationProfile{
			Location: to.Ptr("East US"),
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
				Overrides: []interface{}{
					map[string]interface{}{
						"if": map[string]interface{}{
							"equals": "eastus",
							"field":  "$.location",
						},
						"priority": float64(100),
						"then": map[string]interface{}{
							"LogAnalytics/Enable":      true,
							"LogAnalytics/Reprovision": true,
							"LogAnalytics/Workspace":   "/subscriptions/abc/resourceGroups/xyz/providers/Microsoft.La/Workspaces/eastus",
						},
					},
					map[string]interface{}{
						"if": map[string]interface{}{
							"equals": "centralcanada",
							"field":  "$.location",
						},
						"priority": float64(200),
						"then": map[string]interface{}{
							"LogAnalytics/Enable":      true,
							"LogAnalytics/Reprovision": true,
							"LogAnalytics/Workspace":   "/subscriptions/abc/resourceGroups/xyz/providers/Microsoft.La/Workspaces/centralcanada",
						},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fautomanage%2Farmautomanage%2Fv0.5.0/sdk/resourcemanager/automanage/armautomanage/README.md) on how to add the SDK to your project and authenticate.
