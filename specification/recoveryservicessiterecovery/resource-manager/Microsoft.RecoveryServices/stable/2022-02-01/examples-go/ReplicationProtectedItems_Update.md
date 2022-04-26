Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservicessiterecovery%2Fv0.4.0/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/README.md) on how to add the SDK to your project and authenticate.

```go
package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectedItems_Update.json
func ExampleReplicationProtectedItemsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		"<replicated-protected-item-name>",
		armrecoveryservicessiterecovery.UpdateReplicationProtectedItemInput{
			Properties: &armrecoveryservicessiterecovery.UpdateReplicationProtectedItemInputProperties{
				LicenseType: to.Ptr(armrecoveryservicessiterecovery.LicenseTypeWindowsServer),
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.HyperVReplicaAzureUpdateReplicationProtectedItemInput{
					InstanceType: to.Ptr("<instance-type>"),
				},
				RecoveryAzureVMName:            to.Ptr("<recovery-azure-vmname>"),
				RecoveryAzureVMSize:            to.Ptr("<recovery-azure-vmsize>"),
				SelectedRecoveryAzureNetworkID: to.Ptr("<selected-recovery-azure-network-id>"),
				VMNics: []*armrecoveryservicessiterecovery.VMNicInputDetails{
					{
						IPConfigs: []*armrecoveryservicessiterecovery.IPConfigInputDetails{
							{
								IPConfigName:            to.Ptr("<ipconfig-name>"),
								IsPrimary:               to.Ptr(true),
								RecoveryStaticIPAddress: to.Ptr("<recovery-static-ipaddress>"),
								RecoverySubnetName:      to.Ptr("<recovery-subnet-name>"),
							}},
						NicID:         to.Ptr("<nic-id>"),
						SelectionType: to.Ptr("<selection-type>"),
					}},
			},
		},
		&armrecoveryservicessiterecovery.ReplicationProtectedItemsClientBeginUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
