package armtestbase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/TestSummariesList.json
func ExampleTestSummariesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtestbase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTestSummariesClient().NewListPager("contoso-rg1", "contoso-testBaseAccount1", nil)
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
		// page.TestSummaryListResult = armtestbase.TestSummaryListResult{
		// 	Value: []*armtestbase.TestSummaryResource{
		// 		{
		// 			Name: to.Ptr("contoso-package1-38960b32-3541-4cf1-8ccc-fd22774395cc"),
		// 			Type: to.Ptr("Microsoft.TestBase/testBaseAccounts/testSummaries"),
		// 			ID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg1/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount1/testSummaries/contoso-package1-38960b32-3541-4cf1-8ccc-fd22774395cc"),
		// 			Properties: &armtestbase.TestSummaryProperties{
		// 				ApplicationName: to.Ptr("contoso-package1"),
		// 				ApplicationVersion: to.Ptr("1.0.0"),
		// 				ExecutionStatus: to.Ptr(armtestbase.ExecutionStatusSucceeded),
		// 				FeatureUpdatesTestSummary: &armtestbase.OSUpdatesTestSummary{
		// 					ExecutionStatus: to.Ptr(armtestbase.ExecutionStatusSucceeded),
		// 					Grade: to.Ptr(armtestbase.GradePass),
		// 					OSUpdateTestSummaries: []*armtestbase.OSUpdateTestSummary{
		// 						{
		// 							BuildVersion: to.Ptr("513"),
		// 							ExecutionStatus: to.Ptr(armtestbase.ExecutionStatusSucceeded),
		// 							Grade: to.Ptr(armtestbase.GradePass),
		// 							OSName: to.Ptr("Windows 10 1909"),
		// 							ReleaseName: to.Ptr("2020.12.B"),
		// 							TestRunTime: to.Ptr("00:43:21"),
		// 							TestStatus: to.Ptr(armtestbase.TestStatusCompleted),
		// 							TestType: to.Ptr("OutOfBoxTest"),
		// 						},
		// 						{
		// 							BuildVersion: to.Ptr("765"),
		// 							ExecutionStatus: to.Ptr(armtestbase.ExecutionStatusSucceeded),
		// 							Grade: to.Ptr(armtestbase.GradePass),
		// 							OSName: to.Ptr("Windows 10 1903"),
		// 							ReleaseName: to.Ptr("2020.11.B"),
		// 							TestRunTime: to.Ptr("00:13:28"),
		// 							TestStatus: to.Ptr(armtestbase.TestStatusCompleted),
		// 							TestType: to.Ptr("OutOfBoxTest"),
		// 						},
		// 						{
		// 							BuildVersion: to.Ptr("313"),
		// 							ExecutionStatus: to.Ptr(armtestbase.ExecutionStatusSucceeded),
		// 							Grade: to.Ptr(armtestbase.GradePass),
		// 							OSName: to.Ptr("Windows 10 1809"),
		// 							ReleaseName: to.Ptr("2020.11.B"),
		// 							TestRunTime: to.Ptr("00:42:08"),
		// 							TestStatus: to.Ptr(armtestbase.TestStatusCompleted),
		// 							TestType: to.Ptr("OutOfBoxTest"),
		// 					}},
		// 					TestRunTime: to.Ptr("00:33:21"),
		// 					TestStatus: to.Ptr(armtestbase.TestStatusCompleted),
		// 				},
		// 				Grade: to.Ptr(armtestbase.GradePass),
		// 				PackageID: to.Ptr("57199102-9738-42e0-9fec-db7709d62a71"),
		// 				SecurityUpdatesTestSummary: &armtestbase.OSUpdatesTestSummary{
		// 					ExecutionStatus: to.Ptr(armtestbase.ExecutionStatusSucceeded),
		// 					Grade: to.Ptr(armtestbase.GradePass),
		// 					OSUpdateTestSummaries: []*armtestbase.OSUpdateTestSummary{
		// 						{
		// 							BuildVersion: to.Ptr("513"),
		// 							ExecutionStatus: to.Ptr(armtestbase.ExecutionStatusSucceeded),
		// 							Grade: to.Ptr(armtestbase.GradePass),
		// 							OSName: to.Ptr("Windows 10 1909"),
		// 							ReleaseName: to.Ptr("2020.12.B"),
		// 							TestRunTime: to.Ptr("00:43:21"),
		// 							TestStatus: to.Ptr(armtestbase.TestStatusCompleted),
		// 							TestType: to.Ptr("OutOfBoxTest"),
		// 						},
		// 						{
		// 							BuildVersion: to.Ptr("765"),
		// 							ExecutionStatus: to.Ptr(armtestbase.ExecutionStatusSucceeded),
		// 							Grade: to.Ptr(armtestbase.GradePass),
		// 							OSName: to.Ptr("Windows 10 1903"),
		// 							ReleaseName: to.Ptr("2020.11.B"),
		// 							TestRunTime: to.Ptr("00:13:28"),
		// 							TestStatus: to.Ptr(armtestbase.TestStatusCompleted),
		// 							TestType: to.Ptr("OutOfBoxTest"),
		// 						},
		// 						{
		// 							BuildVersion: to.Ptr("313"),
		// 							ExecutionStatus: to.Ptr(armtestbase.ExecutionStatusSucceeded),
		// 							Grade: to.Ptr(armtestbase.GradePass),
		// 							OSName: to.Ptr("Windows 10 1809"),
		// 							ReleaseName: to.Ptr("2020.11.B"),
		// 							TestRunTime: to.Ptr("00:42:08"),
		// 							TestStatus: to.Ptr(armtestbase.TestStatusCompleted),
		// 							TestType: to.Ptr("OutOfBoxTest"),
		// 					}},
		// 					TestRunTime: to.Ptr("00:43:21"),
		// 					TestStatus: to.Ptr(armtestbase.TestStatusCompleted),
		// 				},
		// 				TestRunTime: to.Ptr("00:43:21"),
		// 				TestStatus: to.Ptr(armtestbase.TestStatusCompleted),
		// 				TestSummaryID: to.Ptr("38960b32-3541-4cf1-8ccc-fd22774395cc"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("contoso-package2-096bffb5-5d3d-4305-a66a-953372ed6e88"),
		// 			Type: to.Ptr("Microsoft.TestBase/testBaseAccounts/testSummaries"),
		// 			ID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg1/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount1/testSummaries/contoso-package2-096bffb5-5d3d-4305-a66a-953372ed6e88"),
		// 			Properties: &armtestbase.TestSummaryProperties{
		// 				ApplicationName: to.Ptr("contoso-package2"),
		// 				ApplicationVersion: to.Ptr("1.0.0"),
		// 				ExecutionStatus: to.Ptr(armtestbase.ExecutionStatusSucceeded),
		// 				FeatureUpdatesTestSummary: &armtestbase.OSUpdatesTestSummary{
		// 					ExecutionStatus: to.Ptr(armtestbase.ExecutionStatusSucceeded),
		// 					Grade: to.Ptr(armtestbase.GradePass),
		// 					OSUpdateTestSummaries: []*armtestbase.OSUpdateTestSummary{
		// 						{
		// 							BuildVersion: to.Ptr("513"),
		// 							ExecutionStatus: to.Ptr(armtestbase.ExecutionStatusSucceeded),
		// 							Grade: to.Ptr(armtestbase.GradePass),
		// 							OSName: to.Ptr("Windows 10 1909"),
		// 							ReleaseName: to.Ptr("2020.12.B"),
		// 							TestRunTime: to.Ptr("00:43:21"),
		// 							TestStatus: to.Ptr(armtestbase.TestStatusCompleted),
		// 							TestType: to.Ptr("FunctionalTest"),
		// 						},
		// 						{
		// 							BuildVersion: to.Ptr("765"),
		// 							ExecutionStatus: to.Ptr(armtestbase.ExecutionStatusSucceeded),
		// 							Grade: to.Ptr(armtestbase.GradePass),
		// 							OSName: to.Ptr("Windows 10 1903"),
		// 							ReleaseName: to.Ptr("2020.11.B"),
		// 							TestRunTime: to.Ptr("00:13:28"),
		// 							TestStatus: to.Ptr(armtestbase.TestStatusCompleted),
		// 							TestType: to.Ptr("FunctionalTest"),
		// 						},
		// 						{
		// 							BuildVersion: to.Ptr("313"),
		// 							ExecutionStatus: to.Ptr(armtestbase.ExecutionStatusSucceeded),
		// 							Grade: to.Ptr(armtestbase.GradePass),
		// 							OSName: to.Ptr("Windows 10 1809"),
		// 							ReleaseName: to.Ptr("2020.11.B"),
		// 							TestRunTime: to.Ptr("00:42:08"),
		// 							TestStatus: to.Ptr(armtestbase.TestStatusCompleted),
		// 							TestType: to.Ptr("FunctionalTest"),
		// 					}},
		// 					TestRunTime: to.Ptr("00:33:21"),
		// 					TestStatus: to.Ptr(armtestbase.TestStatusCompleted),
		// 				},
		// 				Grade: to.Ptr(armtestbase.GradePass),
		// 				PackageID: to.Ptr("b5ed1bcc-e74c-40d8-82f2-1773f616f93e"),
		// 				SecurityUpdatesTestSummary: &armtestbase.OSUpdatesTestSummary{
		// 					ExecutionStatus: to.Ptr(armtestbase.ExecutionStatusSucceeded),
		// 					Grade: to.Ptr(armtestbase.GradePass),
		// 					OSUpdateTestSummaries: []*armtestbase.OSUpdateTestSummary{
		// 						{
		// 							BuildVersion: to.Ptr("513"),
		// 							ExecutionStatus: to.Ptr(armtestbase.ExecutionStatusSucceeded),
		// 							Grade: to.Ptr(armtestbase.GradePass),
		// 							OSName: to.Ptr("Windows 10 1909"),
		// 							ReleaseName: to.Ptr("2020.12.B"),
		// 							TestRunTime: to.Ptr("00:43:21"),
		// 							TestStatus: to.Ptr(armtestbase.TestStatusCompleted),
		// 							TestType: to.Ptr("FunctionalTest"),
		// 						},
		// 						{
		// 							BuildVersion: to.Ptr("765"),
		// 							ExecutionStatus: to.Ptr(armtestbase.ExecutionStatusSucceeded),
		// 							Grade: to.Ptr(armtestbase.GradePass),
		// 							OSName: to.Ptr("Windows 10 1903"),
		// 							ReleaseName: to.Ptr("2020.11.B"),
		// 							TestRunTime: to.Ptr("00:13:28"),
		// 							TestStatus: to.Ptr(armtestbase.TestStatusCompleted),
		// 							TestType: to.Ptr("FunctionalTest"),
		// 						},
		// 						{
		// 							BuildVersion: to.Ptr("313"),
		// 							ExecutionStatus: to.Ptr(armtestbase.ExecutionStatusSucceeded),
		// 							Grade: to.Ptr(armtestbase.GradePass),
		// 							OSName: to.Ptr("Windows 10 1809"),
		// 							ReleaseName: to.Ptr("2020.11.B"),
		// 							TestRunTime: to.Ptr("00:42:08"),
		// 							TestStatus: to.Ptr(armtestbase.TestStatusCompleted),
		// 							TestType: to.Ptr("FunctionalTest"),
		// 					}},
		// 					TestRunTime: to.Ptr("00:43:21"),
		// 					TestStatus: to.Ptr(armtestbase.TestStatusCompleted),
		// 				},
		// 				TestRunTime: to.Ptr("00:43:21"),
		// 				TestStatus: to.Ptr(armtestbase.TestStatusCompleted),
		// 				TestSummaryID: to.Ptr("096bffb5-5d3d-4305-a66a-953372ed6e88"),
		// 			},
		// 	}},
		// }
	}
}
