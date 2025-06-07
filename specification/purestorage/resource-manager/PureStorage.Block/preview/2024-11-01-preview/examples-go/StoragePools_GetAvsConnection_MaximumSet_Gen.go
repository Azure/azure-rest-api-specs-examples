package armpurestorageblock_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purestorageblock/armpurestorageblock"
)

// Generated from example definition: 2024-11-01-preview/StoragePools_GetAvsConnection_MaximumSet_Gen.json
func ExampleStoragePoolsClient_GetAvsConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurestorageblock.NewClientFactory("BC47D6CC-AA80-4374-86F8-19D94EC70666", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStoragePoolsClient().GetAvsConnection(ctx, "rgpurestorage", "storagePoolname", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpurestorageblock.StoragePoolsClientGetAvsConnectionResponse{
	// 	AvsConnection: &armpurestorageblock.AvsConnection{
	// 		ServiceInitializationCompleted: to.Ptr(true),
	// 		ServiceInitializationHandleEnc: to.Ptr("kvidtocmjciflvtwql"),
	// 		ServiceInitializationHandle: &armpurestorageblock.ServiceInitializationHandle{
	// 			ClusterResourceID: to.Ptr("aqgnpie"),
	// 			ServiceAccountUsername: to.Ptr("xcu"),
	// 		},
	// 	},
	// }
}
