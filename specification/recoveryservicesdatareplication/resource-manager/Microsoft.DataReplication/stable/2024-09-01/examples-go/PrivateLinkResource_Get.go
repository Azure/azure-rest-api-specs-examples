package armrecoveryservicesdatareplication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
)

// Generated from example definition: 2024-09-01/PrivateLinkResource_Get.json
func ExamplePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("930CEC23-4430-4513-B855-DBA237E2F3BF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().Get(ctx, "rgswagger_2024-09-01", "4", "d", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrecoveryservicesdatareplication.PrivateLinkResourcesClientGetResponse{
	// 	PrivateLinkResource: &armrecoveryservicesdatareplication.PrivateLinkResource{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataReplication/replicationVaults/vault1/privateLinkResources/link1"),
	// 		Name: to.Ptr("u"),
	// 		Type: to.Ptr("ydabrhxuyafq"),
	// 		Properties: &armrecoveryservicesdatareplication.PrivateLinkResourceProperties{
	// 			GroupID: to.Ptr("dzkqscbkzs"),
	// 			RequiredMembers: []*string{
	// 				to.Ptr("irjjsneakjewcsigcocfdjvfad"),
	// 			},
	// 			RequiredZoneNames: []*string{
	// 				to.Ptr("yp"),
	// 			},
	// 		},
	// 	},
	// }
}
