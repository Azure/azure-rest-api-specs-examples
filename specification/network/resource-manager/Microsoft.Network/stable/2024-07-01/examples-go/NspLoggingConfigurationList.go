package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c58fa033619b12c7cfa8a0ec5a9bf03bb18869ab/specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NspLoggingConfigurationList.json
func ExampleSecurityPerimeterLoggingConfigurationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSecurityPerimeterLoggingConfigurationsClient().NewListPager("rg1", "nsp1", nil)
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
		// page.NspLoggingConfigurationListResult = armnetwork.NspLoggingConfigurationListResult{
		// 	Value: []*armnetwork.NspLoggingConfiguration{
		// 		{
		// 			Name: to.Ptr("instance"),
		// 			Type: to.Ptr("Microsoft.Network/networkSecurityPerimeters/loggingConfigurations"),
		// 			ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp1/loggingConfigurations/instance"),
		// 			SystemData: &armnetwork.SecurityPerimeterSystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-07T18:07:36.344Z"); return t}()),
		// 				CreatedBy: to.Ptr("user"),
		// 				CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-07T18:07:36.344Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user"),
		// 				LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 			},
		// 			Properties: &armnetwork.NspLoggingConfigurationProperties{
		// 				EnabledLogCategories: []*string{
		// 					to.Ptr("NspPublicInboundPerimeterRulesDenied"),
		// 					to.Ptr("NspPublicOutboundPerimeterRulesDenied")},
		// 					Version: to.Ptr("0"),
		// 				},
		// 		}},
		// 	}
	}
}
