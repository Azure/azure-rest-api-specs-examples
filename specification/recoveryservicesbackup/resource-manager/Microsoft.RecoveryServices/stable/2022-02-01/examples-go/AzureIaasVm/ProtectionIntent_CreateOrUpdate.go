package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/AzureIaasVm/ProtectionIntent_CreateOrUpdate.json
func ExampleProtectionIntentClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicesbackup.NewProtectionIntentClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"myVault",
		"myRG",
		"Azure",
		"vm;iaasvmcontainerv2;chamsrgtest;chamscandel",
		armrecoveryservicesbackup.ProtectionIntentResource{
			Properties: &armrecoveryservicesbackup.AzureResourceProtectionIntent{
				PolicyID:                 to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.RecoveryServices/vaults/myVault/backupPolicies/myPolicy"),
				ProtectionIntentItemType: to.Ptr(armrecoveryservicesbackup.ProtectionIntentItemTypeAzureResourceItem),
				SourceResourceID:         to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/chamsrgtest/providers/Microsoft.Compute/virtualMachines/chamscandel"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
