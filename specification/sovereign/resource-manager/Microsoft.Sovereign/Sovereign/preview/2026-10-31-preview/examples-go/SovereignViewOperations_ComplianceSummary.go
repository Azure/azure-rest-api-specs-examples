package armsovereign_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sovereign/armsovereign"
)

// Generated from example definition: 2026-10-31-preview/SovereignViewOperations_ComplianceSummary.json
func ExampleViewOperationsClient_ComplianceSummary() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsovereign.NewClientFactory("22b62f96-c8d5-4131-97f8-f820a5a3c55f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewViewOperationsClient().ComplianceSummary(ctx, "ExampleResourceGroup", "ExampleSV", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsovereign.ViewOperationsClientComplianceSummaryResponse{
	// 	ViewComplianceSummaryResult: armsovereign.ViewComplianceSummaryResult{
	// 		ResourceID: to.Ptr("/subscriptions/22b62f96-c8d5-4131-97f8-f820a5a3c55f/resourceGroups/ExampleResourceGroup/providers/Microsoft.Sovereign/sovereignViews/ExampleSV"),
	// 		ReportStatus: &armsovereign.ReportStatus{
	// 			State: to.Ptr(armsovereign.ReportStateUpToDate),
	// 			NextAttempt: to.Ptr(time.Date(2026, time.August, 1, 6, 20, 21, 0, time.UTC)),
	// 		},
	// 		LastEvaluated: to.Ptr(time.Date(2026, time.March, 15, 12, 30, 0, 0, time.UTC)),
	// 		OverallSummary: &armsovereign.ViewComplianceOverallSummary{
	// 			TotalManagementGroups: to.Ptr[int32](1),
	// 			TotalSubscriptions: to.Ptr[int32](2),
	// 			TotalResourceGroups: to.Ptr[int32](3),
	// 			TotalResources: to.Ptr[int32](11),
	// 			CompliantCount: to.Ptr[int32](5),
	// 			NonCompliantCount: to.Ptr[int32](2),
	// 			ConstrainedCount: to.Ptr[int32](1),
	// 			ConflictCount: to.Ptr[int32](1),
	// 			NotApplicableCount: to.Ptr[int32](1),
	// 			NotEvaluatedCount: to.Ptr[int32](1),
	// 			OthersCount: to.Ptr[int32](0),
	// 		},
	// 		DimensionSummaries: []*armsovereign.ViewComplianceDimensionSummary{
	// 			{
	// 				Dimension: to.Ptr(armsovereign.ComplianceDimensionEncryptionAtRest),
	// 				TotalResources: to.Ptr[int32](10),
	// 				CompliantCount: to.Ptr[int32](5),
	// 				NonCompliantCount: to.Ptr[int32](2),
	// 				ConstrainedCount: to.Ptr[int32](0),
	// 				ConflictCount: to.Ptr[int32](1),
	// 				NotApplicableCount: to.Ptr[int32](1),
	// 				NotEvaluatedCount: to.Ptr[int32](1),
	// 				OthersCount: to.Ptr[int32](0),
	// 			},
	// 			{
	// 				Dimension: to.Ptr(armsovereign.ComplianceDimensionBusinessContinuity),
	// 				TotalResources: to.Ptr[int32](10),
	// 				CompliantCount: to.Ptr[int32](6),
	// 				NonCompliantCount: to.Ptr[int32](1),
	// 				ConstrainedCount: to.Ptr[int32](1),
	// 				ConflictCount: to.Ptr[int32](0),
	// 				NotApplicableCount: to.Ptr[int32](1),
	// 				NotEvaluatedCount: to.Ptr[int32](1),
	// 				OthersCount: to.Ptr[int32](0),
	// 			},
	// 		},
	// 	},
	// }
}
