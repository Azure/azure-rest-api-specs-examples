package armfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/fabric/armfabric"
)

// Generated from example definition: D:/w/Azure/azure-sdk-for-go/sdk/resourcemanager/fabric/armfabric/TempTypeSpecFiles/Microsoft.Fabric.Management/examples/2023-11-01/FabricCapacities_Get.json
func ExampleCapacitiesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfabric.NewClientFactory("548B7FB7-3B2A-4F46-BB02-66473F1FC22C", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCapacitiesClient().Get(ctx, "TestRG", "azsdktest", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armfabric.CapacitiesClientGetResponse{
	// 	Capacity: &armfabric.Capacity{
	// 		Properties: &armfabric.CapacityProperties{
	// 			ProvisioningState: to.Ptr(armfabric.ProvisioningStateSucceeded),
	// 			State: to.Ptr(armfabric.ResourceStateActive),
	// 			Administration: &armfabric.CapacityAdministration{
	// 				Members: []*string{
	// 					to.Ptr("azsdktest@microsoft.com"),
	// 					to.Ptr("azsdktest2@microsoft.com"),
	// 				},
	// 			},
	// 		},
	// 		SKU: &armfabric.RpSKU{
	// 			Name: to.Ptr("F2"),
	// 			Tier: to.Ptr(armfabric.RpSKUTierFabric),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("West Central US"),
	// 		ID: to.Ptr("/subscriptions/548B7FB7-3B2A-4F46-BB02-66473F1FC22C/resourceGroups/TestRG/providers/Microsoft.Fabric/capacities/azsdktest"),
	// 		Name: to.Ptr("azsdktest"),
	// 		Type: to.Ptr("Microsoft.Fabric/capacities"),
	// 	},
	// }
}
