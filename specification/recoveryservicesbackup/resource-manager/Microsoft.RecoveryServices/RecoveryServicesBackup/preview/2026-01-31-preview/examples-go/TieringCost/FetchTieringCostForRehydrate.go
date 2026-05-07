package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v5"
)

// Generated from example definition: 2026-01-31-preview/TieringCost/FetchTieringCostForRehydrate.json
func ExampleFetchTieringCostClient_BeginPost_getTheRehydrationCostForRecoveryPoint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFetchTieringCostClient().BeginPost(ctx, "netsdktestrg", "testVault", &armrecoveryservicesbackup.FetchTieringCostInfoForRehydrationRequest{
		ContainerName:       to.Ptr("IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1"),
		ObjectType:          to.Ptr("FetchTieringCostInfoForRehydrationRequest"),
		ProtectedItemName:   to.Ptr("VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1"),
		RecoveryPointID:     to.Ptr("1222343434"),
		RehydrationPriority: to.Ptr(armrecoveryservicesbackup.RehydrationPriorityHigh),
		SourceTierType:      to.Ptr(armrecoveryservicesbackup.RecoveryPointTierTypeArchivedRP),
		TargetTierType:      to.Ptr(armrecoveryservicesbackup.RecoveryPointTierTypeHardenedRP),
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
	// res = armrecoveryservicesbackup.FetchTieringCostClientPostResponse{
	// 	TieringCostRehydrationInfo: &armrecoveryservicesbackup.TieringCostRehydrationInfo{
	// 		ObjectType: to.Ptr("TieringCostRehydrationInfo"),
	// 		RehydrationSizeInBytes: to.Ptr[int64](1204000),
	// 		RetailRehydrationCostPerGBPerMonth: to.Ptr[float64](0.02),
	// 	},
	// }
}
