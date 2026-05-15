package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v3"
)

// Generated from example definition: 2025-03-01/settingsList.json
func ExampleSettingsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSettingsClient().List(ctx, "subscriptions/00000000-0000-0000-0000-000000000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcostmanagement.SettingsClientListResponse{
	// 	SettingsListResult: armcostmanagement.SettingsListResult{
	// 		Value: []armcostmanagement.SettingClassification{
	// 			&armcostmanagement.TagInheritanceSetting{
	// 				Name: to.Ptr("taginheritance"),
	// 				Type: to.Ptr("Microsoft.CostManagement/Settings"),
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.CostManagement/settings/taginheritance"),
	// 				Kind: to.Ptr(armcostmanagement.SettingsKindTaginheritance),
	// 				Properties: &armcostmanagement.TagInheritanceProperties{
	// 					PreferContainerTags: to.Ptr(false),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
