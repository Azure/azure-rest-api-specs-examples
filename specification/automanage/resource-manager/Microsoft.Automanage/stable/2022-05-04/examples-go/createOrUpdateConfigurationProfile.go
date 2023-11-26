package armautomanage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automanage/armautomanage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/2dcad6d6e9a96882eb6d317e7500a94be007a9c6/specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/createOrUpdateConfigurationProfile.json
func ExampleConfigurationProfilesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomanage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConfigurationProfilesClient().CreateOrUpdate(ctx, "customConfigurationProfile", "myResourceGroupName", armautomanage.ConfigurationProfile{
		Location: to.Ptr("East US"),
		Tags: map[string]*string{
			"Organization": to.Ptr("Administration"),
		},
		Properties: &armautomanage.ConfigurationProfileProperties{
			Configuration: map[string]any{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConfigurationProfile = armautomanage.ConfigurationProfile{
	// 	Name: to.Ptr("customConfigurationProfile"),
	// 	Type: to.Ptr("Microsoft.Automanage/configurationProfiles"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.Automanage/configurationProfiles/customConfigurationProfile"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 		"Organization": to.Ptr("Administration"),
	// 	},
	// 	Properties: &armautomanage.ConfigurationProfileProperties{
	// 		Configuration: map[string]any{
	// 			"Antimalware/Enable": false,
	// 			"AzureSecurityCenter/Enable": true,
	// 			"Backup/Enable": false,
	// 			"BootDiagnostics/Enable": true,
	// 			"ChangeTrackingAndInventory/Enable": true,
	// 			"GuestConfiguration/Enable": true,
	// 			"LogAnalytics/Enable": true,
	// 			"UpdateManagement/Enable": true,
	// 			"VMInsights/Enable": true,
	// 		},
	// 	},
	// 	SystemData: &armautomanage.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.107Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1@outlook.com"),
	// 		CreatedByType: to.Ptr(armautomanage.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-04T02:03:01.197Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2@outlook.com"),
	// 		LastModifiedByType: to.Ptr(armautomanage.CreatedByTypeUser),
	// 	},
	// }
}
