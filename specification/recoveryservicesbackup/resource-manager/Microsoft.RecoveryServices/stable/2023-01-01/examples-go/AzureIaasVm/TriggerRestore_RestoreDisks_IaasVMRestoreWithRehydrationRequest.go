package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3751704f5318f1175875c94b66af769db917f2d3/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-01-01/examples/AzureIaasVm/TriggerRestore_RestoreDisks_IaasVMRestoreWithRehydrationRequest.json
func ExampleRestoresClient_BeginTrigger_restoreDisksWithIaasVmRestoreWithRehydrationRequest() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicesbackup.NewRestoresClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginTrigger(ctx, "testVault", "netsdktestrg", "Azure", "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1", "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1", "348916168024334", armrecoveryservicesbackup.RestoreRequestResource{
		Properties: &armrecoveryservicesbackup.IaasVMRestoreWithRehydrationRequest{
			ObjectType:            to.Ptr("IaasVMRestoreWithRehydrationRequest"),
			CreateNewCloudService: to.Ptr(true),
			EncryptionDetails: &armrecoveryservicesbackup.EncryptionDetails{
				EncryptionEnabled: to.Ptr(false),
			},
			OriginalStorageAccountOption: to.Ptr(false),
			RecoveryPointID:              to.Ptr("348916168024334"),
			RecoveryType:                 to.Ptr(armrecoveryservicesbackup.RecoveryTypeRestoreDisks),
			Region:                       to.Ptr("southeastasia"),
			SourceResourceID:             to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1"),
			StorageAccountID:             to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testingRg/providers/Microsoft.Storage/storageAccounts/testAccount"),
			RecoveryPointRehydrationInfo: &armrecoveryservicesbackup.RecoveryPointRehydrationInfo{
				RehydrationPriority:          to.Ptr(armrecoveryservicesbackup.RehydrationPriorityStandard),
				RehydrationRetentionDuration: to.Ptr("P7D"),
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
