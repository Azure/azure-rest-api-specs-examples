Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservicesbackup%2Fv0.3.0/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```go
package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/AzureIaasVm/Provision_Ilr.json
func ExampleItemLevelRecoveryConnectionsClient_Provision() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewItemLevelRecoveryConnectionsClient("<subscription-id>", cred, nil)
	_, err = client.Provision(ctx,
		"<vault-name>",
		"<resource-group-name>",
		"<fabric-name>",
		"<container-name>",
		"<protected-item-name>",
		"<recovery-point-id>",
		armrecoveryservicesbackup.ILRRequestResource{
			Properties: &armrecoveryservicesbackup.IaasVMILRRegistrationRequest{
				ObjectType:                to.StringPtr("<object-type>"),
				InitiatorName:             to.StringPtr("<initiator-name>"),
				RecoveryPointID:           to.StringPtr("<recovery-point-id>"),
				RenewExistingRegistration: to.BoolPtr(true),
				VirtualMachineID:          to.StringPtr("<virtual-machine-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
