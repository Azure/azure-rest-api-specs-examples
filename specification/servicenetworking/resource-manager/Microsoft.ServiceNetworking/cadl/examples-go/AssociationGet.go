package armservicenetworking_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicenetworking/armservicenetworking"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/71121282e39bccae590462648e77bca283df6d2b/specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/AssociationGet.json
func ExampleAssociationsInterfaceClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicenetworking.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssociationsInterfaceClient().Get(ctx, "rg1", "TC1", "associatedvnet-2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Association = armservicenetworking.Association{
	// 	Name: to.Ptr("associatedvnet-2"),
	// 	Type: to.Ptr("Microsoft.ServiceNetworking/trafficControllers/associations"),
	// 	ID: to.Ptr("resourceUriAsString"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armservicenetworking.AssociationProperties{
	// 		AssociationType: to.Ptr("subnets"),
	// 		ProvisioningState: to.Ptr(armservicenetworking.ProvisioningStateSucceeded),
	// 		Subnet: &armservicenetworking.AssociationSubnet{
	// 			ID: to.Ptr("subnetFullRef"),
	// 		},
	// 	},
	// }
}
