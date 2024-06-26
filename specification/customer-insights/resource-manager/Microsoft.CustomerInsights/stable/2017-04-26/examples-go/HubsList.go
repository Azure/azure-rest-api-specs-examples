package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/HubsList.json
func ExampleHubsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewHubsClient().NewListPager(nil)
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
		// page.HubListResult = armcustomerinsights.HubListResult{
		// 	Value: []*armcustomerinsights.Hub{
		// 		{
		// 			Name: to.Ptr("azSdkTestHub"),
		// 			Type: to.Ptr("Microsoft.CustomerInsights/hubs"),
		// 			ID: to.Ptr("/subscriptions/c909e979-ef71-4def-a970-bc7c154db8c5/resourceGroups/TestHubRG/providers/Microsoft.CustomerInsights/hubs/azSdkTestHub"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armcustomerinsights.HubPropertiesFormat{
		// 				APIEndpoint: to.Ptr("https://azSdkTestHub.dfd.projectuic-int.net"),
		// 				HubBillingInfo: &armcustomerinsights.HubBillingInfoFormat{
		// 					MaxUnits: to.Ptr[int32](1),
		// 					MinUnits: to.Ptr[int32](1),
		// 					SKUName: to.Ptr("B0"),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				TenantFeatures: to.Ptr[int32](0),
		// 				WebEndpoint: to.Ptr("https://azSdkTestHub.dfdapps.projectuic-int.net"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testHub1058"),
		// 			Type: to.Ptr("Microsoft.CustomerInsights/hubs"),
		// 			ID: to.Ptr("/subscriptions/c909e979-ef71-4def-a970-bc7c154db8c5/resourceGroups/TestHubRG/providers/Microsoft.CustomerInsights/hubs/testHub1058"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armcustomerinsights.HubPropertiesFormat{
		// 				APIEndpoint: to.Ptr("https://testHub1058.dfd.projectuic-int.net"),
		// 				HubBillingInfo: &armcustomerinsights.HubBillingInfoFormat{
		// 					MaxUnits: to.Ptr[int32](5),
		// 					MinUnits: to.Ptr[int32](1),
		// 					SKUName: to.Ptr("B0"),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				TenantFeatures: to.Ptr[int32](0),
		// 				WebEndpoint: to.Ptr("https://testHub1058.dfdapps.projectuic-int.net"),
		// 			},
		// 	}},
		// }
	}
}
