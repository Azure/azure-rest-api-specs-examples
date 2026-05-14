package armcostmanagement_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v3"
)

// Generated from example definition: 2025-03-01/BenefitUtilizationSummaries/Async/GenerateBenefitUtilizationSummariesReportByBillingAccount.json
func ExampleGenerateBenefitUtilizationSummariesReportClient_BeginGenerateByBillingAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGenerateBenefitUtilizationSummariesReportClient().BeginGenerateByBillingAccount(ctx, "8099099", armcostmanagement.BenefitUtilizationSummariesRequest{
		EndDate:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-31T00:00:00Z"); return t }()),
		Grain:     to.Ptr(armcostmanagement.GrainDaily),
		Kind:      to.Ptr(armcostmanagement.BenefitKindReservation),
		StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-06-01T00:00:00Z"); return t }()),
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
	// res = armcostmanagement.GenerateBenefitUtilizationSummariesReportClientGenerateByBillingAccountResponse{
	// 	BenefitUtilizationSummariesOperationStatus: armcostmanagement.BenefitUtilizationSummariesOperationStatus{
	// 		Input: &armcostmanagement.BenefitUtilizationSummariesRequest{
	// 			BillingAccountID: to.Ptr("8099099"),
	// 			EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-31T00:00:00Z"); return t}()),
	// 			Grain: to.Ptr(armcostmanagement.GrainDaily),
	// 			Kind: to.Ptr(armcostmanagement.BenefitKindReservation),
	// 			StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-06-01T00:00:00Z"); return t}()),
	// 		},
	// 		Properties: &armcostmanagement.AsyncOperationStatusProperties{
	// 			ReportURL: to.Ptr(armcostmanagement.BenefitUtilizationSummaryReportSchema("https://storage.blob.core.windows.net/details/20220611/00000000-0000-0000-0000-000000000000?sv=2016-05-31&sr=b&sig=jep8HT2aphfUkyERRZa5LRfd9RPzjXbzB%2F9TNiQ")),
	// 			SecondaryReportURL: to.Ptr(armcostmanagement.BenefitUtilizationSummaryReportSchema("https://storage-secondary.blob.core.windows.net/details/20220611/00000000-0000-0000-0000-000000000000?sv=2016-05-31&sr=b&sig=jep8HT2aphfUkyERRZa5LRfd9RPzjXbzB%2F9TNiQ")),
	// 			ValidUntil: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-06-12T02:56:55.5021869Z"); return t}()),
	// 		},
	// 		Status: to.Ptr(armcostmanagement.OperationStatusType("Complete")),
	// 	},
	// }
}
