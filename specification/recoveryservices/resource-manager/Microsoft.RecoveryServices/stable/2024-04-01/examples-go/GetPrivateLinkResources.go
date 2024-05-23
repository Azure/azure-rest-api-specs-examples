package armrecoveryservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/GetPrivateLinkResources.json
func ExamplePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().Get(ctx, "petesting", "pemsi-ecy-rsv2", "backupResource", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResource = armrecoveryservices.PrivateLinkResource{
	// 	Name: to.Ptr("backupResource"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/Vaults/privateLinkResources"),
	// 	ID: to.Ptr("/subscriptions/6c48fa17-39c7-45f1-90ac-47a587128ace/resourceGroups/petesting/providers/Microsoft.RecoveryServices/Vaults/pemsi-ecy-rsv2/privateLinkResources/backupResource"),
	// 	Properties: &armrecoveryservices.PrivateLinkResourceProperties{
	// 		GroupID: to.Ptr("AzureBackup"),
	// 		RequiredMembers: []*string{
	// 			to.Ptr("backup-fab1"),
	// 			to.Ptr("backup-rec2"),
	// 			to.Ptr("backup-prot1"),
	// 			to.Ptr("backup-ecs1"),
	// 			to.Ptr("backup-tel1"),
	// 			to.Ptr("backup-wbcm1"),
	// 			to.Ptr("backup-fc1"),
	// 			to.Ptr("backup-id1")},
	// 			RequiredZoneNames: []*string{
	// 				to.Ptr("privatelink.ecy.backup.windowsazure.com"),
	// 				to.Ptr("privatelink.queue.core.windows.net"),
	// 				to.Ptr("privatelink.blob.core.windows.net")},
	// 			},
	// 		}
}
