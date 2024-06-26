package armautomanage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automanage/armautomanage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/2dcad6d6e9a96882eb6d317e7500a94be007a9c6/specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listConfigurationProfileVersions.json
func ExampleConfigurationProfilesVersionsClient_NewListChildResourcesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomanage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConfigurationProfilesVersionsClient().NewListChildResourcesPager("customConfigurationProfile", "myResourceGroupName", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ConfigurationProfileList = armautomanage.ConfigurationProfileList{
		// 	Value: []*armautomanage.ConfigurationProfile{
		// 		{
		// 			Name: to.Ptr("customConfigurationProfile/version1"),
		// 			Type: to.Ptr("Microsoft.Automanage/ConfigurationProfiles/versions"),
		// 			ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.Automanage/configurationProfiles/customConfigurationProfile/versions/versions1"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 				"Organization": to.Ptr("Administration"),
		// 			},
		// 			Properties: &armautomanage.ConfigurationProfileProperties{
		// 				Configuration: map[string]any{
		// 					"Antimalware/Enable": false,
		// 					"AzureSecurityCenter/Enable": true,
		// 					"Backup/Enable": false,
		// 					"BootDiagnostics/Enable": true,
		// 					"ChangeTrackingAndInventory/Enable": true,
		// 					"GuestConfiguration/Enable": true,
		// 					"LogAnalytics/Enable": true,
		// 					"UpdateManagement/Enable": true,
		// 					"VMInsights/Enable": true,
		// 				},
		// 			},
		// 			SystemData: &armautomanage.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.107Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1@outlook.com"),
		// 				CreatedByType: to.Ptr(armautomanage.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-04T02:03:01.197Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2@outlook.com"),
		// 				LastModifiedByType: to.Ptr(armautomanage.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("customConfigurationProfile/version2"),
		// 			Type: to.Ptr("Microsoft.Automanage/ConfigurationProfiles/versions"),
		// 			ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.Automanage/configurationProfiles/customConfigurationProfile/versions/version2"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 				"Organization": to.Ptr("Administration"),
		// 			},
		// 			Properties: &armautomanage.ConfigurationProfileProperties{
		// 				Configuration: map[string]any{
		// 					"Antimalware/Enable": false,
		// 					"AzureSecurityCenter/Enable": true,
		// 					"Backup/Enable": false,
		// 					"BootDiagnostics/Enable": true,
		// 					"ChangeTrackingAndInventory/Enable": true,
		// 					"GuestConfiguration/Enable": true,
		// 					"LogAnalytics/Enable": true,
		// 					"UpdateManagement/Enable": true,
		// 					"VMInsights/Enable": true,
		// 				},
		// 			},
		// 			SystemData: &armautomanage.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.107Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1@outlook.com"),
		// 				CreatedByType: to.Ptr(armautomanage.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-04T02:03:01.197Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2@outlook.com"),
		// 				LastModifiedByType: to.Ptr(armautomanage.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
