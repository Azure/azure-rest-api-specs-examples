package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/ResourceGuardProxyCRUD/UnlockDeleteResourceGuardProxy.json
func ExampleDppResourceGuardProxyClient_UnlockDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewDppResourceGuardProxyClient("5e13b949-1218-4d18-8b99-7e12155ec4f7", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.UnlockDelete(ctx, "SampleResourceGroup", "sampleVault", "swaggerExample", armdataprotection.UnlockDeleteRequest{
		ResourceGuardOperationRequests: []*string{
			to.Ptr("/subscriptions/f9e67185-f313-4e79-aa71-6458d429369d/resourceGroups/ResourceGuardSecurityAdminRG/providers/Microsoft.DataProtection/resourceGuards/ResourceGuardTestResource/deleteBackupInstanceRequests/default")},
		ResourceToBeDeleted: to.Ptr("/subscriptions/5e13b949-1218-4d18-8b99-7e12155ec4f7/resourceGroups/SampleResourceGroup/providers/Microsoft.DataProtection/backupVaults/sampleVault/backupInstances/TestBI9779f4de"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
