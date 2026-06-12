package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: 2025-07-01-preview/settings/GetAllSettings.json
func ExampleProductSettingsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProductSettingsClient().NewListPager("myRg", "myWorkspace", nil)
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
		// page = armsecurityinsights.ProductSettingsClientListResponse{
		// 	SettingList: armsecurityinsights.SettingList{
		// 		Value: []armsecurityinsights.SettingsClassification{
		// 			&armsecurityinsights.EyesOn{
		// 				Name: to.Ptr("EyesOn"),
		// 				Type: to.Ptr("Microsoft.SecurityInsights/settings"),
		// 				ID: to.Ptr("/subscriptions/bd794837-4d29-4647-9105-6339bfdb4e6a/resourceGroups/mms-eus/providers/Microsoft.OperationalInsights/workspaces/avdvirInt/providers/Microsoft.SecurityInsights/settings/EyesOn"),
		// 				Kind: to.Ptr(armsecurityinsights.SettingKindEyesOn),
		// 				Properties: &armsecurityinsights.EyesOnSettingsProperties{
		// 					IsEnabled: to.Ptr(true),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
