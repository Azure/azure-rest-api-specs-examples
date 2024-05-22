package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/TieringCost/FetchTieringCostForRehydrate.json
func ExampleFetchTieringCostClient_BeginPost_getTheRehydrationCostForRecoveryPoint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFetchTieringCostClient().BeginPost(ctx, "netsdktestrg", "testVault", &armrecoveryservicesbackup.FetchTieringCostInfoForRehydrationRequest{
		ObjectType:          to.Ptr("FetchTieringCostInfoForRehydrationRequest"),
		SourceTierType:      to.Ptr(armrecoveryservicesbackup.RecoveryPointTierTypeArchivedRP),
		TargetTierType:      to.Ptr(armrecoveryservicesbackup.RecoveryPointTierTypeHardenedRP),
		ContainerName:       to.Ptr("IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1"),
		ProtectedItemName:   to.Ptr("VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1"),
		RecoveryPointID:     to.Ptr("1222343434"),
		RehydrationPriority: to.Ptr(armrecoveryservicesbackup.RehydrationPriorityHigh),
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
	// res = armrecoveryservicesbackup.FetchTieringCostClientPostResponse{
	// 	                            TieringCostInfoClassification: &armrecoveryservicesbackup.TieringCostRehydrationInfo{
	// 		ObjectType: to.Ptr("TieringCostRehydrationInfo"),
	// 		RehydrationSizeInBytes: to.Ptr[int64](1204000),
	// 		RetailRehydrationCostPerGBPerMonth: to.Ptr[float64](0.02),
	// 	},
	// 	                        }
}
