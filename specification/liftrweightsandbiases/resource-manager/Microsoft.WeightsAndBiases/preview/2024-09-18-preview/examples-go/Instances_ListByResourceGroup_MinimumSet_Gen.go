package armweightsandbiases_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/weightsandbiases/armweightsandbiases"
)

// Generated from example definition: 2024-09-18-preview/Instances_ListByResourceGroup_MinimumSet_Gen.json
func ExampleInstancesClient_NewListByResourceGroupPager_instancesListByResourceGroupGeneratedByMinimumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armweightsandbiases.NewClientFactory("0BCB047F-CB71-4DFE-B0BD-88519F411C2F", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInstancesClient().NewListByResourceGroupPager("rgopenapi", nil)
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
		// page = armweightsandbiases.InstancesClientListByResourceGroupResponse{
		// 	InstanceResourceListResult: armweightsandbiases.InstanceResourceListResult{
		// 		Value: []*armweightsandbiases.InstanceResource{
		// 			{
		// 				ID: to.Ptr("/subscriptions/0BCB047F-CB71-4DFE-B0BD-88519F411C2F/resourceGroups/myResourceGroup/providers/Microsoft.WeightsAndBiases/instance/myinstance"),
		// 				Location: to.Ptr("pudewmshbcvbt"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
