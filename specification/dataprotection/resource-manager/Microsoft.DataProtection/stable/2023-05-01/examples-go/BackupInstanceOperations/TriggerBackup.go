package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-05-01/examples/BackupInstanceOperations/TriggerBackup.json
func ExampleBackupInstancesClient_BeginAdhocBackup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.OperationJobExtendedInfo = armdataprotection.OperationJobExtendedInfo{
	// 	ObjectType: to.Ptr("OperationJobExtendedInfo"),
	// 	JobID: to.Ptr("c60cb49-63e8-4b21-b9bd-26277b3fdfae"),
	// }
}
