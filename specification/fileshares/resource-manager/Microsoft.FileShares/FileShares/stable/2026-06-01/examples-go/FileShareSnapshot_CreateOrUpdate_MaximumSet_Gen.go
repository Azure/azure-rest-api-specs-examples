package armfileshares_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/fileshares/armfileshares"
)

// Generated from example definition: 2026-06-01/FileShareSnapshot_CreateOrUpdate_MaximumSet_Gen.json
func ExampleFileShareSnapshotsClient_BeginCreateOrUpdateFileShareSnapshot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfileshares.NewClientFactory("0681745E-3F9F-4966-80E6-69624A3B29F2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFileShareSnapshotsClient().BeginCreateOrUpdateFileShareSnapshot(ctx, "rgfileshares", "fileshare", "testfilesharesnapshot", armfileshares.FileShareSnapshot{
		Properties: &armfileshares.FileShareSnapshotProperties{
			InitiatorID: to.Ptr("backup-vault-001"),
			Metadata: map[string]*string{
				"key9372": to.Ptr("jtc"),
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
