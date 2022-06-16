package armautomanage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automanage/armautomanage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automanage/resource-manager/Microsoft.Automanage/preview/2021-04-30-preview/examples/updateConfigurationProfileVersion.json
func ExampleConfigurationProfilesVersionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armautomanage.NewConfigurationProfilesVersionsClient("mySubscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"customConfigurationProfile",
		"version1",
		"myResourceGroupName",
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
	}
	// TODO: use response item
	_ = res
}
