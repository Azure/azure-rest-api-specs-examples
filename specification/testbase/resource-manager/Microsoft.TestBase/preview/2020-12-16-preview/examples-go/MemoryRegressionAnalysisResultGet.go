package armtestbase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/MemoryRegressionAnalysisResultGet.json
func ExampleAnalysisResultsClient_Get_memoryRegressionAnalysisResultGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtestbase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAnalysisResultsClient().Get(ctx, "contoso-rg1", "contoso-testBaseAccount1", "contoso-package2", "Windows-10-1909-Test-Id", armtestbase.AnalysisResultNameMemoryRegression, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AnalysisResultSingletonResource = armtestbase.AnalysisResultSingletonResource{
	// 	Name: to.Ptr("memoryRegression"),
	// 	Type: to.Ptr("Microsoft.TestBase/testBaseAccounts/packages/testResults/analysisResults"),
	// 	ID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg1/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount1/packages/contoso-package2/testResults/Windows-10-1909-Test-Id/analysisResults/memoryRegression"),
	// 	Properties: &armtestbase.MemoryRegressionResultSingletonResourceProperties{
	// 		AnalysisResultType: to.Ptr(armtestbase.AnalysisResultTypeMemoryRegression),
	// 		Grade: to.Ptr(armtestbase.GradePass),
	// 		MemoryRegressionResults: []*armtestbase.RegressionResult{
	// 			{
	// 				Diff: to.Ptr[float64](0.1),
	// 				FileName: to.Ptr("testApp.exe"),
	// 				Grade: to.Ptr(armtestbase.GradePass),
	// 				IsRegressed: to.Ptr(false),
	// 				Details: to.Ptr("Some details of testApp"),
	// 			},
	// 			{
	// 				Diff: to.Ptr[float64](0.15),
	// 				FileName: to.Ptr("dependencies.exe"),
	// 				Grade: to.Ptr(armtestbase.GradePass),
	// 				IsRegressed: to.Ptr(false),
	// 				Details: to.Ptr("Some details of dependencies"),
	// 		}},
	// 	},
	// }
}
