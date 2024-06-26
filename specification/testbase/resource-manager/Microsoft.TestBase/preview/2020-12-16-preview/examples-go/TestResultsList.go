package armtestbase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/TestResultsList.json
func ExampleTestResultsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtestbase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTestResultsClient().NewListPager("contoso-rg1", "contoso-testBaseAccount1", "contoso-package2", armtestbase.OsUpdateTypeSecurityUpdate, &armtestbase.TestResultsClientListOptions{Filter: to.Ptr("osName eq 'Windows 10 2004' and releaseName eq '2020.11B'")})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.TestResultListResult = armtestbase.TestResultListResult{
		// 	Value: []*armtestbase.TestResultResource{
		// 		{
		// 			Name: to.Ptr("Windows-10-2004-8e8e3200-284b-4a99-8c30-cb46c54d4c1e"),
		// 			Type: to.Ptr("Microsoft.TestBase/testBaseAccounts/packages/testResults"),
		// 			ID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg1/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount1/packages/contoso-package2/testResults/Windows-10-2004-8e8e3200-284b-4a99-8c30-cb46c54d4c1e"),
		// 			Properties: &armtestbase.TestResultProperties{
		// 				AnalysisSummaries: []*armtestbase.TestResultAnalysisSummary{
		// 					{
		// 						Name: to.Ptr("Memory Regression Analysis Result"),
		// 						AnalysisStatus: to.Ptr(armtestbase.AnalysisStatusSucceeded),
		// 						Grade: to.Ptr(armtestbase.GradePass),
		// 					},
		// 					{
		// 						Name: to.Ptr("CPU Regression Analysis Result"),
		// 						AnalysisStatus: to.Ptr(armtestbase.AnalysisStatusSucceeded),
		// 						Grade: to.Ptr(armtestbase.GradePass),
		// 					},
		// 					{
		// 						Name: to.Ptr("Memory Utilization Analysis Result"),
		// 						AnalysisStatus: to.Ptr(armtestbase.AnalysisStatusSucceeded),
		// 						Grade: to.Ptr(armtestbase.GradePass),
		// 					},
		// 					{
		// 						Name: to.Ptr("CPU Utilization Analysis Result"),
		// 						AnalysisStatus: to.Ptr(armtestbase.AnalysisStatusSucceeded),
		// 						Grade: to.Ptr(armtestbase.GradePass),
		// 				}},
		// 				ApplicationName: to.Ptr("contoso-package2"),
		// 				ApplicationVersion: to.Ptr("1.0.0"),
		// 				BaselineTestResultID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg1/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount1/packages/contoso-package2/testResults/anotherId"),
		// 				BuildRevision: to.Ptr("478"),
		// 				ExecutionStatus: to.Ptr(armtestbase.ExecutionStatusSucceeded),
		// 				Grade: to.Ptr(armtestbase.GradePass),
		// 				IsDownloadDataAvailable: to.Ptr(true),
		// 				KbNumber: to.Ptr("KB1984839"),
		// 				OSName: to.Ptr("Windows 10 2004"),
		// 				PackageID: to.Ptr("b5ed1bcc-e74c-40d8-82f2-1773f616f93e"),
		// 				PackageVersion: to.Ptr("3.0.1"),
		// 				ReleaseName: to.Ptr("2020.11B"),
		// 				ReleaseVersionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-28T17:30:00.000Z"); return t}()),
		// 				TestRunTime: to.Ptr("00:21:30"),
		// 				TestStatus: to.Ptr(armtestbase.TestStatusCompleted),
		// 				TestType: to.Ptr("Out of box test"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Windows-10-2004-d19baae0-9ab4-432b-a752-878343fa4481"),
		// 			Type: to.Ptr("Microsoft.TestBase/testBaseAccounts/packages/testResults"),
		// 			ID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg1/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount1/packages/contoso-package2/testResults/Windows-10-2004-d19baae0-9ab4-432b-a752-878343fa4481"),
		// 			Properties: &armtestbase.TestResultProperties{
		// 				AnalysisSummaries: []*armtestbase.TestResultAnalysisSummary{
		// 					{
		// 						Name: to.Ptr("Memory Regression Analysis Result"),
		// 						AnalysisStatus: to.Ptr(armtestbase.AnalysisStatusNotAvailable),
		// 						Grade: to.Ptr(armtestbase.GradeNotAvailable),
		// 					},
		// 					{
		// 						Name: to.Ptr("CPU Regression Analysis Result"),
		// 						AnalysisStatus: to.Ptr(armtestbase.AnalysisStatusNotAvailable),
		// 						Grade: to.Ptr(armtestbase.GradeNotAvailable),
		// 					},
		// 					{
		// 						Name: to.Ptr("Memory Utilization Analysis Result"),
		// 						AnalysisStatus: to.Ptr(armtestbase.AnalysisStatusNotAvailable),
		// 						Grade: to.Ptr(armtestbase.GradeNotAvailable),
		// 					},
		// 					{
		// 						Name: to.Ptr("CPU Utilization Analysis Result"),
		// 						AnalysisStatus: to.Ptr(armtestbase.AnalysisStatusNotAvailable),
		// 						Grade: to.Ptr(armtestbase.GradeNotAvailable),
		// 				}},
		// 				ApplicationName: to.Ptr("contoso-package2"),
		// 				ApplicationVersion: to.Ptr("1.0.0"),
		// 				BaselineTestResultID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg1/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount1/packages/contoso-package2/testResults/anotherId"),
		// 				BuildRevision: to.Ptr("309"),
		// 				ExecutionStatus: to.Ptr(armtestbase.ExecutionStatusFailed),
		// 				Grade: to.Ptr(armtestbase.GradeFail),
		// 				IsDownloadDataAvailable: to.Ptr(false),
		// 				KbNumber: to.Ptr("KB1784831"),
		// 				OSName: to.Ptr("Windows 10 2004"),
		// 				PackageID: to.Ptr("b5ed1bcc-e74c-40d8-82f2-1773f616f93e"),
		// 				PackageVersion: to.Ptr("3.0.1"),
		// 				ReleaseName: to.Ptr("2020.11B"),
		// 				ReleaseVersionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-28T17:30:00.000Z"); return t}()),
		// 				TestRunTime: to.Ptr("01:00:00"),
		// 				TestStatus: to.Ptr(armtestbase.TestStatusTestFailure),
		// 				TestType: to.Ptr("Out of box test"),
		// 			},
		// 	}},
		// }
	}
}
