package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/ResourceGuardProxyCRUD/UnlockDeleteResourceGuardProxy.json
func ExampleResourceGuardProxyClient_UnlockDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewResourceGuardProxyClient("<subscription-id>", cred, nil)
	_, err = client.UnlockDelete(ctx,
		"<vault-name>",
		"<resource-group-name>",
		"<resource-guard-proxy-name>",
		armrecoveryservicesbackup.UnlockDeleteRequest{
			ResourceGuardOperationRequests: []*string{
				to.StringPtr("/subscriptions/c999d45b-944f-418c-a0d8-c3fcfd1802c8/resourceGroups/vaultguardRGNew/providers/Microsoft.DataProtection/resourceGuards/VaultGuardTestNew/deleteProtectedItemRequests/default")},
			ResourceToBeDeleted: to.StringPtr("<resource-to-be-deleted>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
