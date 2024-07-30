package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a651ba25cda4eec698a3a4e35f867ecc2681d126/specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/stable/2024-04-01/examples/managedUnsupportedVMSizesGet_example.json
func ExampleManagedUnsupportedVMSizesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedUnsupportedVMSizesClient().Get(ctx, "eastus", "Standard_B1ls1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedVMSize = armservicefabricmanagedclusters.ManagedVMSize{
	// 	Name: to.Ptr("Standard_B1ls1"),
	// 	Type: to.Ptr("Microsoft.ServiceFabric/locations/managedVMSizes"),
	// 	ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ServiceFabric/locations/eastus/managedVMSizes/Standard_B1ls1"),
	// 	Properties: &armservicefabricmanagedclusters.VMSize{
	// 		Size: to.Ptr("Standard_B1ls1"),
	// 	},
	// }
}
