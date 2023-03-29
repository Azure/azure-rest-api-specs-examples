package armdelegatednetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/delegatednetwork/armdelegatednetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/patchDelegatedSubnet.json
func ExampleDelegatedSubnetServiceClient_BeginPatchDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdelegatednetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDelegatedSubnetServiceClient().BeginPatchDetails(ctx, "TestRG", "delegated1", armdelegatednetwork.ResourceUpdateParameters{
		Tags: map[string]*string{
			"key": to.Ptr("value"),
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
	// res.DelegatedSubnet = armdelegatednetwork.DelegatedSubnet{
	// 	Name: to.Ptr("delegated1"),
	// 	Type: to.Ptr("Microsoft.DelegatedNetwork/delegatedSubnets"),
	// 	ID: to.Ptr("/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/TestRG/providers/Microsoft.DelegatedNetwork/delegatedSubnets/delegated1"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Properties: &armdelegatednetwork.DelegatedSubnetProperties{
	// 		ControllerDetails: &armdelegatednetwork.ControllerDetails{
	// 			ID: to.Ptr("/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/TestRG/providers/Microsoft.DelegatedNetwork/controller/dnctestcontroller"),
	// 		},
	// 		ProvisioningState: to.Ptr(armdelegatednetwork.DelegatedSubnetStateSucceeded),
	// 		ResourceGUID: to.Ptr("5a82cbcf-e8ea-4175-ac2b-ad36a73f9801"),
	// 		SubnetDetails: &armdelegatednetwork.SubnetDetails{
	// 			ID: to.Ptr("/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/TestRG/providers/Microsoft.Network/virtualNetworks/testvnet/subnets/testsubnet"),
	// 		},
	// 	},
	// }
}
