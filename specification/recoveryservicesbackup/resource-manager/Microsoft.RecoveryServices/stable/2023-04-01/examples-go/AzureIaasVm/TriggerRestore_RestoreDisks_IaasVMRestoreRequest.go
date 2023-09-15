package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a4ddec441435d1ef766c4f160eda658a69cc5dc2/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/AzureIaasVm/TriggerRestore_RestoreDisks_IaasVMRestoreRequest.json
func ExampleRestoresClient_BeginTrigger_restoreDisksWithIaasVmRestoreRequest() {
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
			CreateNewCloudService: to.Ptr(true),
			EncryptionDetails: &armrecoveryservicesbackup.EncryptionDetails{
				EncryptionEnabled: to.Ptr(false),
			},
			IdentityInfo: &armrecoveryservicesbackup.IdentityInfo{
				IsSystemAssignedIdentity:  to.Ptr(false),
				ManagedIdentityResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/asmaskarRG1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/asmaskartestmsi"),
			},
			OriginalStorageAccountOption: to.Ptr(false),
			RecoveryPointID:              to.Ptr("348916168024334"),
			RecoveryType:                 to.Ptr(armrecoveryservicesbackup.RecoveryTypeRestoreDisks),
			Region:                       to.Ptr("southeastasia"),
			SourceResourceID:             to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1"),
			StorageAccountID:             to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testingRg/providers/Microsoft.Storage/storageAccounts/testAccount"),
			TargetDiskNetworkAccessSettings: &armrecoveryservicesbackup.TargetDiskNetworkAccessSettings{
				TargetDiskAccessID:            to.Ptr("/subscriptions/e7a191f5-713c-4bdb-b5e4-cf3dd90230ef/resourceGroups/arpja/providers/Microsoft.Compute/diskAccesses/arpja-diskaccess-ccy"),
				TargetDiskNetworkAccessOption: to.Ptr(armrecoveryservicesbackup.TargetDiskNetworkAccessOptionEnablePrivateAccessForAllDisks),
			},
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
