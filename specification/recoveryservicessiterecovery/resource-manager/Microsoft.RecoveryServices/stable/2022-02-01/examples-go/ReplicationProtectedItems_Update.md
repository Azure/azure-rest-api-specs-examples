Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservicessiterecovery%2Fv1.0.0/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/README.md) on how to add the SDK to your project and authenticate.

```go
package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectedItems_Update.json
func ExampleReplicationProtectedItemsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("vault1",
		"resourceGroupPS1",
		"c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"cloud1",
		"cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
		"f8491e4f-817a-40dd-a90c-af773978c75b",
		armrecoveryservicessiterecovery.UpdateReplicationProtectedItemInput{
			Properties: &armrecoveryservicessiterecovery.UpdateReplicationProtectedItemInputProperties{
				LicenseType: to.Ptr(armrecoveryservicessiterecovery.LicenseTypeWindowsServer),
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.HyperVReplicaAzureUpdateReplicationProtectedItemInput{
					InstanceType: to.Ptr("HyperVReplicaAzure"),
				},
				RecoveryAzureVMName:            to.Ptr("vm1"),
				RecoveryAzureVMSize:            to.Ptr("Basic_A0"),
				SelectedRecoveryAzureNetworkID: to.Ptr("/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/siterecoveryProd1/providers/Microsoft.Network/virtualNetworks/vnetavrai"),
				VMNics: []*armrecoveryservicessiterecovery.VMNicInputDetails{
					{
						IPConfigs: []*armrecoveryservicessiterecovery.IPConfigInputDetails{
							{
								IPConfigName:            to.Ptr("ipconfig1"),
								IsPrimary:               to.Ptr(true),
								RecoveryStaticIPAddress: to.Ptr("10.0.2.46"),
								RecoverySubnetName:      to.Ptr("subnet1"),
							}},
						NicID:         to.Ptr("TWljcm9zb2Z0OkY4NDkxRTRGLTgxN0EtNDBERC1BOTBDLUFGNzczOTc4Qzc1Qlw3NjAwMzMxRS03NDk4LTQ0QTQtQjdDNy0xQjY1NkJDREQ1MkQ="),
						SelectionType: to.Ptr("SelectedByUser"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
