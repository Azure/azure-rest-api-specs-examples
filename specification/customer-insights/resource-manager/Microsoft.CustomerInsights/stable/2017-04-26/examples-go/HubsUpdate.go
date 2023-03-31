package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/HubsUpdate.json
func ExampleHubsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewHubsClient().Update(ctx, "TestHubRG", "sdkTestHub", armcustomerinsights.Hub{
		Location: to.Ptr("West US"),
		Properties: &armcustomerinsights.HubPropertiesFormat{
			HubBillingInfo: &armcustomerinsights.HubBillingInfoFormat{
				MaxUnits: to.Ptr[int32](5),
				MinUnits: to.Ptr[int32](1),
				SKUName:  to.Ptr("B0"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Hub = armcustomerinsights.Hub{
	// 	Name: to.Ptr("testHub2839"),
	// 	Type: to.Ptr("Microsoft.CustomerInsights/hubs"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/TestHubRG/providers/Microsoft.CustomerInsights/hubs/testHub2839"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armcustomerinsights.HubPropertiesFormat{
	// 		APIEndpoint: to.Ptr("https://testHub2839.dfd.projectuic-int.net"),
	// 		HubBillingInfo: &armcustomerinsights.HubBillingInfoFormat{
	// 			MaxUnits: to.Ptr[int32](5),
	// 			MinUnits: to.Ptr[int32](1),
	// 			SKUName: to.Ptr("B0"),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		TenantFeatures: to.Ptr[int32](0),
	// 		WebEndpoint: to.Ptr("https://testHub2839.dfdapps.projectuic-int.net"),
	// 	},
	// }
}
