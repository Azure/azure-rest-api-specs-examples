package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5700885250d8f685a17293e930d98d1c1d72f401/specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-11-01/examples/Volumes_LatestRestoreStatus.json
func ExampleBackupsClient_GetVolumeLatestRestoreStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBackupsClient().GetVolumeLatestRestoreStatus(ctx, "myRG", "account1", "pool1", "volume1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RestoreStatus = armnetapp.RestoreStatus{
	// 	ErrorMessage: to.Ptr(""),
	// 	Healthy: to.Ptr(true),
	// 	MirrorState: to.Ptr(armnetapp.MirrorStateUninitialized),
	// 	RelationshipStatus: to.Ptr(armnetapp.RelationshipStatusIdle),
	// 	TotalTransferBytes: to.Ptr[int64](100000),
	// 	UnhealthyReason: to.Ptr(""),
	// }
}
