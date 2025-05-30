package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/148c3b0b44f7789ced94859992493fafd0072f83/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/AzureIaasVm/Provision_Ilr.json
func ExampleItemLevelRecoveryConnectionsClient_Provision() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewItemLevelRecoveryConnectionsClient().Provision(ctx, "PySDKBackupTestRsVault", "PythonSDKBackupTestRg", "Azure", "iaasvmcontainer;iaasvmcontainerv2;pysdktestrg;pysdktestv2vm1", "vm;iaasvmcontainerv2;pysdktestrg;pysdktestv2vm1", "1", armrecoveryservicesbackup.ILRRequestResource{
		Properties: &armrecoveryservicesbackup.IaasVMILRRegistrationRequest{
			ObjectType:                to.Ptr("IaasVMILRRegistrationRequest"),
			InitiatorName:             to.Ptr("Hello World"),
			RecoveryPointID:           to.Ptr("38823086363464"),
			RenewExistingRegistration: to.Ptr(true),
			VirtualMachineID:          to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/pysdktestrg/providers/Microsoft.Compute/virtualMachines/pysdktestv2vm1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
