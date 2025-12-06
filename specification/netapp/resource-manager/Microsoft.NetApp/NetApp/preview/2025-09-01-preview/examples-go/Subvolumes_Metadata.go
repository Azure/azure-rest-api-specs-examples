package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/Subvolumes_Metadata.json
func ExampleSubvolumesClient_BeginGetMetadata() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSubvolumesClient().BeginGetMetadata(ctx, "myRG", "account1", "pool1", "volume1", "subvolume1", nil)
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
	// res = armnetapp.SubvolumesClientGetMetadataResponse{
	// 	SubvolumeModel: &armnetapp.SubvolumeModel{
	// 		Name: to.Ptr("account1/pool1/volume1/subvolume1/metadata"),
	// 		Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/subvolumes/metadata"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/volume1/subvolumes/subvolume1/metadata"),
	// 		Properties: &armnetapp.SubvolumeModelProperties{
	// 			Path: to.Ptr("/pathToSubvol"),
	// 			AccessedTimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-15T13:23:33Z"); return t}()),
	// 			BytesUsed: to.Ptr[int64](5),
	// 			ChangedTimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-15T13:23:33Z"); return t}()),
	// 			CreationTimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-15T13:23:33Z"); return t}()),
	// 			ModifiedTimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-15T13:23:33Z"); return t}()),
	// 			Permissions: to.Ptr("777"),
	// 			Size: to.Ptr[int64](5),
	// 		},
	// 	},
	// }
}
