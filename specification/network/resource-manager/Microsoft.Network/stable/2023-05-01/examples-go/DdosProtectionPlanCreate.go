package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/DdosProtectionPlanCreate.json
func ExampleDdosProtectionPlansClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDdosProtectionPlansClient().BeginCreateOrUpdate(ctx, "rg1", "test-plan", armnetwork.DdosProtectionPlan{
		Location:   to.Ptr("westus"),
		Properties: &armnetwork.DdosProtectionPlanPropertiesFormat{},
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
	// res.DdosProtectionPlan = armnetwork.DdosProtectionPlan{
	// 	Name: to.Ptr("test-plan"),
	// 	Type: to.Ptr("Microsoft.Network/ddosProtectionPlans"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/ddosProtectionPlans/test-plan"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armnetwork.DdosProtectionPlanPropertiesFormat{
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		PublicIPAddresses: []*armnetwork.SubResource{
	// 		},
	// 		ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		VirtualNetworks: []*armnetwork.SubResource{
	// 		},
	// 	},
	// }
}
