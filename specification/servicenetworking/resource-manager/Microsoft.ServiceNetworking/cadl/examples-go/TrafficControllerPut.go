package armservicenetworking_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicenetworking/armservicenetworking"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/71121282e39bccae590462648e77bca283df6d2b/specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/TrafficControllerPut.json
func ExampleTrafficControllerInterfaceClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicenetworking.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTrafficControllerInterfaceClient().BeginCreateOrUpdate(ctx, "rg1", "TC1", armservicenetworking.TrafficController{
		Location: to.Ptr("West US"),
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TrafficController = armservicenetworking.TrafficController{
	// 	Name: to.Ptr("atc"),
	// 	Type: to.Ptr("Microsoft.ServiceNetworking/TrafficController"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ServiceNetworking/trafficController/tc"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 	},
	// 	Properties: &armservicenetworking.TrafficControllerProperties{
	// 		Associations: []*armservicenetworking.ResourceID{
	// 			{
	// 				ID: to.Ptr("/sub1/rg2/tc/tcame/association/as1"),
	// 		}},
	// 		ConfigurationEndpoints: []*string{
	// 			to.Ptr("abc.westus.trafficcontroller.azure.net")},
	// 			Frontends: []*armservicenetworking.ResourceID{
	// 				{
	// 					ID: to.Ptr("/sub1/rg2/tc/tcame/frontends/fe1"),
	// 			}},
	// 			ProvisioningState: to.Ptr(armservicenetworking.ProvisioningStateSucceeded),
	// 		},
	// 	}
}
