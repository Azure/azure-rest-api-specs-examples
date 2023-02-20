package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3035214dea933cd02b1ecfa982c185a572f84b8a/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-01-01/examples/BackupInstanceOperations/TriggerRehydrate.json
func ExampleBackupInstancesClient_BeginTriggerRehydrate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginTriggerRehydrate(ctx, "000pikumar", "PratikPrivatePreviewVault1", "testInstance1", armdataprotection.AzureBackupRehydrationRequest{
		RecoveryPointID:              to.Ptr("hardcodedRP"),
		RehydrationPriority:          to.Ptr(armdataprotection.RehydrationPriorityHigh),
		RehydrationRetentionDuration: to.Ptr("7D"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
