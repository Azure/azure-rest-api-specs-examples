package armrecoveryservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices/v3"
)

// Generated from example definition: 2026-05-01/DeletedVaults_Get.json
func ExampleDeletedVaultsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservices.NewClientFactory("77777777-b0c6-47a2-b37c-d8e65a629c18", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeletedVaultsClient().Get(ctx, "westus", "swaggerExample", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrecoveryservices.DeletedVaultsClientGetResponse{
	// 	DeletedVault: armrecoveryservices.DeletedVault{
	// 		ID: to.Ptr("/subscriptions/77777777-b0c6-47a2-b37c-d8e65a629c18/providers/Microsoft.RecoveryServices/locations/westus/deletedVaults/swaggerExample"),
	// 		Name: to.Ptr("swaggerExample"),
	// 		Type: to.Ptr("Microsoft.RecoveryServices/locations/deletedVaults"),
	// 		Properties: &armrecoveryservices.DeletedVaultProperties{
	// 			VaultID: to.Ptr("/subscriptions/77777777-b0c6-47a2-b37c-d8e65a629c18/resourceGroups/Default-RecoveryServices-ResourceGroup/providers/Microsoft.RecoveryServices/vaults/swaggerExample"),
	// 			VaultDeletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-09-20T09:49:46Z"); return t}()),
	// 			PurgeAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-20T09:49:46Z"); return t}()),
	// 		},
	// 	},
	// }
}
