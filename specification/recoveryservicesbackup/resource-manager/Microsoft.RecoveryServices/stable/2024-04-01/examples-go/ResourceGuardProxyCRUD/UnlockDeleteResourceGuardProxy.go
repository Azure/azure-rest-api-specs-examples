package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/ResourceGuardProxyCRUD/UnlockDeleteResourceGuardProxy.json
func ExampleResourceGuardProxyClient_UnlockDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewResourceGuardProxyClient().UnlockDelete(ctx, "sampleVault", "SampleResourceGroup", "swaggerExample", armrecoveryservicesbackup.UnlockDeleteRequest{
		ResourceGuardOperationRequests: []*string{
			to.Ptr("/subscriptions/c999d45b-944f-418c-a0d8-c3fcfd1802c8/resourceGroups/vaultguardRGNew/providers/Microsoft.DataProtection/resourceGuards/VaultGuardTestNew/deleteProtectedItemRequests/default")},
		ResourceToBeDeleted: to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/gaallarg/providers/Microsoft.RecoveryServices/vaults/MercuryCrrVault/backupFabrics/Azure/protectionContainers/VMAppContainer;compute;crrtestrg;crrtestvm/protectedItems/SQLDataBase;mssqlserver;testdb"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.UnlockDeleteResponse = armrecoveryservicesbackup.UnlockDeleteResponse{
	// 	UnlockDeleteExpiryTime: to.Ptr("2021-02-11T13:12:27.7870742Z"),
	// }
}
