package armresiliencemanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resiliencemanagement/armresiliencemanagement"
)

// Generated from example definition: 2026-08-31-preview/DrillRuns_GenerateReport_MaximumSet_Gen.json
func ExampleDrillRunsClient_BeginGenerateReport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresiliencemanagement.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDrillRunsClient().BeginGenerateReport(ctx, "sampleServiceGroupName", "qmn", "drill1", "ca92602e-53bf-43d2-ae62-d3fc940474b3", nil)
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
	// res = armresiliencemanagement.DrillRunsClientGenerateReportResponse{
	// 	DrillReportSummary: armresiliencemanagement.DrillReportSummary{
	// 		GenerationStatus: to.Ptr(armresiliencemanagement.DrillReportGenerationStatusSucceeded),
	// 		StageStatuses: []*armresiliencemanagement.ReportStageStatus{
	// 			{
	// 				DrillRunStage: to.Ptr(armresiliencemanagement.DrillRunSubtasksFaultInjection),
	// 				GenerationStatus: to.Ptr(armresiliencemanagement.DrillReportGenerationStatusSucceeded),
	// 				LastAttemptTimestamp: to.Ptr(time.Date(2026, time.August, 31, 10, 35, 0, 0, time.UTC)),
	// 			},
	// 		},
	// 		AvailableFormats: []*armresiliencemanagement.DrillReportFormat{
	// 			to.Ptr(armresiliencemanagement.DrillReportFormatHTML),
	// 		},
	// 		LastGeneratedTimestamp: to.Ptr(time.Date(2026, time.August, 31, 10, 35, 0, 0, time.UTC)),
	// 		SchemaVersion: to.Ptr("1.0"),
	// 		FinalizationState: to.Ptr(armresiliencemanagement.DrillReportFinalizationStateNotFinalized),
	// 	},
	// }
}
