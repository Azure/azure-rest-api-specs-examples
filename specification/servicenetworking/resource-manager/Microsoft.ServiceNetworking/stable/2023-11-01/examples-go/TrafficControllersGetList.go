package armservicenetworking_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicenetworking/armservicenetworking"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/stable/2023-11-01/examples/TrafficControllersGetList.json
func ExampleTrafficControllerInterfaceClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicenetworking.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTrafficControllerInterfaceClient().NewListBySubscriptionPager(nil)
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
		// page.TrafficControllerListResult = armservicenetworking.TrafficControllerListResult{
		// 	Value: []*armservicenetworking.TrafficController{
		// 		{
		// 			Name: to.Ptr("tc1"),
		// 			Type: to.Ptr("Microsoft.ServiceNetworking/trafficControllers"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ServiceNetworking/trafficControllers/tc1"),
		// 			Location: to.Ptr("NorthCentralUS"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 			},
		// 			Properties: &armservicenetworking.TrafficControllerProperties{
		// 				Associations: []*armservicenetworking.ResourceID{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ServiceNetworking/trafficControllers/tc1/association/as1"),
		// 				}},
		// 				ConfigurationEndpoints: []*string{
		// 					to.Ptr("abc.trafficcontroller.azure.net")},
		// 					Frontends: []*armservicenetworking.ResourceID{
		// 						{
		// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ServiceNetworking/trafficControllers/tc1/frontends/fe1"),
		// 					}},
		// 					ProvisioningState: to.Ptr(armservicenetworking.ProvisioningStateSucceeded),
		// 				},
		// 		}},
		// 	}
	}
}
