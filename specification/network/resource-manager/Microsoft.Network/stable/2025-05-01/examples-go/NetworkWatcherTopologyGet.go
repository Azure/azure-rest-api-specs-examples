package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkWatcherTopologyGet.json
func ExampleWatchersClient_GetTopology() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWatchersClient().GetTopology(ctx, "rg1", "nw1", armnetwork.TopologyParameters{
		TargetResourceGroupName: to.Ptr("rg2"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Topology = armnetwork.Topology{
	// 	CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-02T19:31:55.946Z"); return t}()),
	// 	ID: to.Ptr("ce592f46-8164-4bf2-ad36-b8e4acf6fb68"),
	// 	LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-27T00:00:13.200Z"); return t}()),
	// 	Resources: []*armnetwork.TopologyResource{
	// 		{
	// 			Name: to.Ptr("MultiTierApp0"),
	// 			Associations: []*armnetwork.TopologyAssociation{
	// 				{
	// 					Name: to.Ptr("appNic0"),
	// 					AssociationType: to.Ptr(armnetwork.AssociationTypeContains),
	// 					ResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/networkInterfaces/appNic0"),
	// 				},
	// 				{
	// 					Name: to.Ptr("appNic10"),
	// 					AssociationType: to.Ptr(armnetwork.AssociationTypeContains),
	// 					ResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/networkInterfaces/appNic10"),
	// 			}},
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/MultiTierApp0"),
	// 			Location: to.Ptr("westus"),
	// 	}},
	// }
}
