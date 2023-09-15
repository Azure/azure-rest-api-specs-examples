package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a4ddec441435d1ef766c4f160eda658a69cc5dc2/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/AzureIaasVm/TriggerRestore_ALR_IaasVMRestoreRequest_IdentityBasedRestoreDetails.json
func ExampleRestoresClient_BeginTrigger_restoreToNewAzureIaasVmWithIaasVmRestoreRequestWithIdentityBasedRestoreDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRestoresClient().BeginTrigger(ctx, "testVault", "netsdktestrg", "Azure", "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1", "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1", "348916168024334", armrecoveryservicesbackup.RestoreRequestResource{
		Properties: &armrecoveryservicesbackup.IaasVMRestoreRequest{
			ObjectType:            to.Ptr("IaasVMRestoreRequest"),
			CreateNewCloudService: to.Ptr(false),
			EncryptionDetails: &armrecoveryservicesbackup.EncryptionDetails{
				EncryptionEnabled: to.Ptr(false),
			},
			IdentityBasedRestoreDetails: &armrecoveryservicesbackup.IdentityBasedRestoreDetails{
				TargetStorageAccountID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.Storage/storageAccounts/testingAccount"),
			},
			IdentityInfo: &armrecoveryservicesbackup.IdentityInfo{
				IsSystemAssignedIdentity: to.Ptr(true),
			},
			OriginalStorageAccountOption: to.Ptr(false),
			RecoveryPointID:              to.Ptr("348916168024334"),
			RecoveryType:                 to.Ptr(armrecoveryservicesbackup.RecoveryTypeAlternateLocation),
			Region:                       to.Ptr("southeastasia"),
			SourceResourceID:             to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1"),
			SubnetID:                     to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.Network/virtualNetworks/testNet/subnets/default"),
			TargetResourceGroupID:        to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg2"),
			TargetVirtualMachineID:       to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg2/providers/Microsoft.Compute/virtualmachines/RSMDALRVM981435"),
			VirtualNetworkID:             to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.Network/virtualNetworks/testNet"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
