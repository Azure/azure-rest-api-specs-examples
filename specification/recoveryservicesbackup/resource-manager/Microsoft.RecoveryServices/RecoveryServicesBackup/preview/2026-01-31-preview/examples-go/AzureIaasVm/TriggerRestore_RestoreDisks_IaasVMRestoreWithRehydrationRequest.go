package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v5"
)

// Generated from example definition: 2026-01-31-preview/AzureIaasVm/TriggerRestore_RestoreDisks_IaasVMRestoreWithRehydrationRequest.json
func ExampleRestoresClient_BeginTrigger_restoreDisksWithIaasVMRestoreWithRehydrationRequest() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRestoresClient().BeginTrigger(ctx, "testVault", "netsdktestrg", "Azure", "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1", "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1", "348916168024334", armrecoveryservicesbackup.RestoreRequestResource{
		Properties: &armrecoveryservicesbackup.IaasVMRestoreWithRehydrationRequest{
			CreateNewCloudService: to.Ptr(true),
			EncryptionDetails: &armrecoveryservicesbackup.EncryptionDetails{
				EncryptionEnabled: to.Ptr(false),
			},
			ObjectType:                   to.Ptr("IaasVMRestoreWithRehydrationRequest"),
			OriginalStorageAccountOption: to.Ptr(false),
			RecoveryPointID:              to.Ptr("348916168024334"),
			RecoveryPointRehydrationInfo: &armrecoveryservicesbackup.RecoveryPointRehydrationInfo{
				RehydrationPriority:          to.Ptr(armrecoveryservicesbackup.RehydrationPriorityStandard),
				RehydrationRetentionDuration: to.Ptr("P7D"),
			},
			RecoveryType:     to.Ptr(armrecoveryservicesbackup.RecoveryTypeRestoreDisks),
			Region:           to.Ptr("southeastasia"),
			SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1"),
			StorageAccountID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testingRg/providers/Microsoft.Storage/storageAccounts/testAccount"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
}
