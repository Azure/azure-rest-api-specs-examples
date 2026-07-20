package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-02/snapshotExamples/Snapshot_UpdateImmutabilityPolicy.json
func ExampleSnapshotsClient_BeginUpdateImmutabilityPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSnapshotsClient().BeginUpdateImmutabilityPolicy(ctx, "myResourceGroup", "mySnapshot", armcompute.ImmutabilityPolicyData{
		ImmutabilityDurationDays: to.Ptr[int32](30),
		Type:                     to.Ptr(armcompute.ImmutabilityPolicyTypeUnlocked),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.SnapshotsClientUpdateImmutabilityPolicyResponse{
	// 	Snapshot: armcompute.Snapshot{
	// 		Properties: &armcompute.SnapshotProperties{
	// 			CreationData: &armcompute.CreationData{
	// 				CreateOption: to.Ptr(armcompute.DiskCreateOptionCopy),
	// 				SourceResourceID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot1"),
	// 			},
	// 			DiskSizeGB: to.Ptr[int32](20),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			ImmutabilityPolicy: &armcompute.ImmutabilityPolicy{
	// 				ImmutabilityDurationDays: to.Ptr[int32](30),
	// 				Type: to.Ptr(armcompute.ImmutabilityPolicyTypeUnlocked),
	// 				PolicyStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-21T12:00:00.0000000+00:00"); return t}()),
	// 				PolicyExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-06-20T12:00:00.0000000+00:00"); return t}()),
	// 				IsPolicyExpired: to.Ptr(false),
	// 			},
	// 		},
	// 		Location: to.Ptr("West US"),
	// 		Name: to.Ptr("mySnapshot"),
	// 	},
	// }
}
