```go
package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/AzureIaasVm/Provision_Ilr.json
func ExampleItemLevelRecoveryConnectionsClient_Provision() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicesbackup.NewItemLevelRecoveryConnectionsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Provision(ctx,
		"PySDKBackupTestRsVault",
		"PythonSDKBackupTestRg",
		"Azure",
		"iaasvmcontainer;iaasvmcontainerv2;pysdktestrg;pysdktestv2vm1",
		"vm;iaasvmcontainerv2;pysdktestrg;pysdktestv2vm1",
		"1",
		armrecoveryservicesbackup.ILRRequestResource{
			Properties: &armrecoveryservicesbackup.IaasVMILRRegistrationRequest{
				ObjectType:                to.Ptr("IaasVMILRRegistrationRequest"),
				InitiatorName:             to.Ptr("Hello World"),
				RecoveryPointID:           to.Ptr("38823086363464"),
				RenewExistingRegistration: to.Ptr(true),
				VirtualMachineID:          to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/pysdktestrg/providers/Microsoft.Compute/virtualMachines/pysdktestv2vm1"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservicesbackup%2Fv1.0.0/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.
