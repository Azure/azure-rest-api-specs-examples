package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/148c3b0b44f7789ced94859992493fafd0072f83/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/TieringCost/FetchTieringCostForProtectedItem.json
func ExampleFetchTieringCostClient_BeginPost_getTheTieringSavingsCostInfoForProtectedItem() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFetchTieringCostClient().BeginPost(ctx, "netsdktestrg", "testVault", &armrecoveryservicesbackup.FetchTieringCostSavingsInfoForProtectedItemRequest{
		ObjectType:        to.Ptr("FetchTieringCostSavingsInfoForProtectedItemRequest"),
		SourceTierType:    to.Ptr(armrecoveryservicesbackup.RecoveryPointTierTypeHardenedRP),
		TargetTierType:    to.Ptr(armrecoveryservicesbackup.RecoveryPointTierTypeArchivedRP),
		ContainerName:     to.Ptr("IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1"),
		ProtectedItemName: to.Ptr("VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1"),
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
	// 	                            TieringCostInfoClassification: &armrecoveryservicesbackup.TieringCostSavingInfo{
	// 		ObjectType: to.Ptr("TieringCostSavingInfo"),
	// 		RetailSourceTierCostPerGBPerMonth: to.Ptr[float64](0.02),
	// 		RetailTargetTierCostPerGBPerMonth: to.Ptr[float64](0.003),
	// 		SourceTierSizeReductionInBytes: to.Ptr[int64](1204000),
	// 		TargetTierSizeIncreaseInBytes: to.Ptr[int64](1892),
	// 	},
	// 	                        }
}
