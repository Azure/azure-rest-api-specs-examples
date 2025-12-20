package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2025-09-01/Datastores_CreateOrUpdate.json
func ExampleDatastoresClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatastoresClient().BeginCreateOrUpdate(ctx, "group1", "cloud1", "cluster1", "datastore1", armavs.Datastore{
		Properties: &armavs.DatastoreProperties{
			NetAppVolume: &armavs.NetAppVolume{
				ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/ResourceGroup1/providers/Microsoft.NetApp/netAppAccounts/NetAppAccount1/capacityPools/CapacityPool1/volumes/NFSVol1"),
			},
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
	// res = armavs.DatastoresClientCreateOrUpdateResponse{
	// 	Datastore: &armavs.Datastore{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/datastores/datastore1"),
	// 		Name: to.Ptr("datastore1"),
	// 		Properties: &armavs.DatastoreProperties{
	// 			NetAppVolume: &armavs.NetAppVolume{
	// 				ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/ResourceGroup1/providers/Microsoft.NetApp/netAppAccounts/NetAppAccount1/capacityPools/CapacityPool1/volumes/NFSVol1"),
	// 			},
	// 			ProvisioningState: to.Ptr(armavs.DatastoreProvisioningStateSucceeded),
	// 			Status: to.Ptr(armavs.DatastoreStatusAccessible),
	// 		},
	// 		Type: to.Ptr("Microsoft.AVS/privateClouds/clusters/datastores"),
	// 	},
	// }
}
