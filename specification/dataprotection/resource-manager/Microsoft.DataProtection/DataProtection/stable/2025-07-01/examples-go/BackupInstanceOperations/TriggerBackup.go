package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v4"
)

// Generated from example definition: 2025-07-01/BackupInstanceOperations/TriggerBackup.json
func ExampleBackupInstancesClient_BeginAdhocBackup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupInstancesClient().BeginAdhocBackup(ctx, "000pikumar", "PratikPrivatePreviewVault1", "testInstance1", armdataprotection.TriggerBackupRequest{
		BackupRuleOptions: &armdataprotection.AdHocBackupRuleOptions{
			RuleName: to.Ptr("BackupWeekly"),
			TriggerOption: &armdataprotection.AdhocBackupTriggerOption{
				RetentionTagOverride: to.Ptr("yearly"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdataprotection.BackupInstancesClientAdhocBackupResponse{
	// 	OperationJobExtendedInfo: &armdataprotection.OperationJobExtendedInfo{
	// 		JobID: to.Ptr("c60cb49-63e8-4b21-b9bd-26277b3fdfae"),
	// 		ObjectType: to.Ptr("OperationJobExtendedInfo"),
	// 	},
	// }
}
