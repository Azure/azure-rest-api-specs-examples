package armoep_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oep/armoep"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/oep/resource-manager/Microsoft.OpenEnergyPlatform/preview/2022-04-04-preview/examples/OepResource_ListByResourceGroup.json
func ExampleEnergyServicesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoep.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEnergyServicesClient().NewListByResourceGroupPager("DummyResourceGroupName", nil)
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
		// page.EnergyServiceList = armoep.EnergyServiceList{
		// 	Value: []*armoep.EnergyService{
		// 		{
		// 			Name: to.Ptr("DummyResourceName"),
		// 			Type: to.Ptr("Microsoft.OEP/oepResource"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000001/resourceGroups/DummyResourceGroupName/providers/Microsoft.OEP/oepResource/DummyResourceName"),
		// 			Location: to.Ptr("WestUS"),
		// 			Properties: &armoep.EnergyServiceProperties{
		// 				AuthAppID: to.Ptr("sample-id"),
		// 				DataPartitionNames: []*armoep.DataPartitionNames{
		// 					{
		// 						Name: to.Ptr("dataPartition1"),
		// 					},
		// 					{
		// 						Name: to.Ptr("dataPartition2"),
		// 				}},
		// 				ProvisioningState: to.Ptr(armoep.ProvisioningStateSucceeded),
		// 			},
		// 			Tags: map[string]*string{
		// 				"additionalProps1": to.Ptr("additional properties"),
		// 			},
		// 	}},
		// }
	}
}
