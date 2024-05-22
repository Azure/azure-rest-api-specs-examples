package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/AzureIaasVm/ProtectionIntent_Validate.json
func ExampleProtectionIntentClient_Validate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProtectionIntentClient().Validate(ctx, "southeastasia", armrecoveryservicesbackup.PreValidateEnableBackupRequest{
		Properties:   to.Ptr(""),
		ResourceID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/arunaupgrade/providers/Microsoft.Compute/VirtualMachines/upgrade1"),
		ResourceType: to.Ptr(armrecoveryservicesbackup.DataSourceTypeVM),
		VaultID:      to.Ptr("/Subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.RecoveryServices/Vaults/myVault"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PreValidateEnableBackupResponse = armrecoveryservicesbackup.PreValidateEnableBackupResponse{
	// 	ContainerName: to.Ptr("iaasvmcontainer;iaasvmcontainerv2;arunaupgrade;upgrade1"),
	// 	ErrorCode: to.Ptr("VirtualMachineAlreadyProtected"),
	// 	ErrorMessage: to.Ptr("Virtual machine with same name and same resource group is already protected. Please select `Disable' choice above for backup and go to backup item corresponding to this VM in the vault"),
	// 	ProtectedItemName: to.Ptr("vm;iaasvmcontainerv2;arunaupgrade;upgrade1"),
	// 	Recommendation: to.Ptr("Please do not enable protection again."),
	// 	Status: to.Ptr(armrecoveryservicesbackup.ValidationStatusFailed),
	// }
}
