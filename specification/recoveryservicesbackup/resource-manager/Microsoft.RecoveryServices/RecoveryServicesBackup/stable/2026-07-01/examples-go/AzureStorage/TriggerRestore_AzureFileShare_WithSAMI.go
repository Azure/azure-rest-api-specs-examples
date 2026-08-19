package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v5"
)

// Generated from example definition: 2026-07-01/AzureStorage/TriggerRestore_AzureFileShare_WithSAMI.json
func ExampleRestoresClient_BeginTrigger_restoreAzureFileShareToOriginalLocationWithManagedIdentity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRestoresClient().BeginTrigger(ctx, "swaggertestvault", "SwaggerTestRg", "Azure", "StorageContainer;Storage;SwaggerTestRg;swaggertestsa", "AzureFileShare;testshare", "932886657837421071", armrecoveryservicesbackup.RestoreRequestResource{
		Properties: &armrecoveryservicesbackup.AzureFileShareRestoreRequest{
			ObjectType:         to.Ptr("AzureFileShareRestoreRequest"),
			RecoveryType:       to.Ptr(armrecoveryservicesbackup.RecoveryTypeOriginalLocation),
			SourceResourceID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/SwaggerTestRg/providers/Microsoft.Storage/storageAccounts/swaggertestsa"),
			CopyOptions:        to.Ptr(armrecoveryservicesbackup.CopyOptionsOverwrite),
			RestoreRequestType: to.Ptr(armrecoveryservicesbackup.RestoreRequestTypeFullShareRestore),
			IdentityInfo: &armrecoveryservicesbackup.IdentityInfo{
				IsSystemAssignedIdentity: to.Ptr(true),
			},
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
