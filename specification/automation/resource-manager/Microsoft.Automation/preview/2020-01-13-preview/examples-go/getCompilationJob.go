package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getCompilationJob.json
func ExampleDscCompilationJobClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDscCompilationJobClient().Get(ctx, "rg", "myAutomationAccount33", "TestCompilationJob", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DscCompilationJob = armautomation.DscCompilationJob{
	// 	Name: to.Ptr("TestCompilationJob"),
	// 	Type: to.Ptr("Microsoft.Automation/AutomationAccounts/compilationjobs"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/compilationjobs/TestCompilationJob"),
	// 	Properties: &armautomation.DscCompilationJobProperties{
	// 		Configuration: &armautomation.DscConfigurationAssociationProperty{
	// 			Name: to.Ptr("SetupServer"),
	// 		},
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:26.903Z"); return t}()),
	// 		JobID: to.Ptr("ce6fe3e3-9db3-4096-a6b4-82bfb4c10a9a"),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:26.903Z"); return t}()),
	// 		LastStatusModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:26.903Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armautomation.JobProvisioningStateSucceeded),
	// 		Status: to.Ptr(armautomation.JobStatusNew),
	// 		StatusDetails: to.Ptr("None"),
	// 	},
	// }
}
