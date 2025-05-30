package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/148c3b0b44f7789ced94859992493fafd0072f83/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/AzureIaasVm/GetBackupStatus.json
func ExampleBackupStatusClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBackupStatusClient().Get(ctx, "southeastasia", armrecoveryservicesbackup.BackupStatusRequest{
		ResourceID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.Compute/VirtualMachines/testVm"),
		ResourceType: to.Ptr(armrecoveryservicesbackup.DataSourceTypeVM),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackupStatusResponse = armrecoveryservicesbackup.BackupStatusResponse{
	// 	ContainerName: to.Ptr("iaasvmcontainer;iaasvmcontainerv2;testRg;testVm"),
	// 	ErrorCode: to.Ptr("Success"),
	// 	ErrorMessage: to.Ptr("ErrorMessage"),
	// 	FabricName: to.Ptr(armrecoveryservicesbackup.FabricNameAzure),
	// 	PolicyName: to.Ptr("myPolicy"),
	// 	ProtectedItemName: to.Ptr("vm;iaasvmcontainerv2;testRg;testVm"),
	// 	ProtectionStatus: to.Ptr(armrecoveryservicesbackup.ProtectionStatusProtected),
	// 	VaultID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.RecoveryServices/Vaults/testVault"),
	// }
}
