package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/639ecfad68419328658bd4cfe7094af4ce472be2/specification/netapp/resource-manager/Microsoft.NetApp/preview/2023-05-01-preview/examples/Volumes_LatestBackupStatus.json
func ExampleBackupsClient_GetLatestStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBackupsClient().GetLatestStatus(ctx, "myRG", "account1", "pool1", "volume1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackupStatus = armnetapp.BackupStatus{
	// 	ErrorMessage: to.Ptr(""),
	// 	Healthy: to.Ptr(true),
	// 	LastTransferSize: to.Ptr[int64](100000),
	// 	LastTransferType: to.Ptr(""),
	// 	MirrorState: to.Ptr(armnetapp.MirrorStateMirrored),
	// 	RelationshipStatus: to.Ptr(armnetapp.RelationshipStatusIdle),
	// 	TotalTransferBytes: to.Ptr[int64](100000),
	// 	UnhealthyReason: to.Ptr(""),
	// }
}
