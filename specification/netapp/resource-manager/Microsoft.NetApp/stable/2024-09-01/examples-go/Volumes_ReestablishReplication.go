package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e6892bbc13d89929cdbe3b39385628543263f80b/specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-09-01/examples/Volumes_ReestablishReplication.json
func ExampleVolumesClient_BeginReestablishReplication() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumesClient().BeginReestablishReplication(ctx, "myRG", "account1", "pool1", "volume1", armnetapp.ReestablishReplicationRequest{
		SourceVolumeID: to.Ptr("/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/resourceGroups/mySourceRG/providers/Microsoft.NetApp/netAppAccounts/sourceAccount1/capacityPools/sourcePool1/volumes/sourceVolume1"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
