package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfigurationMachineRun/getSoftwareUpdateConfigurationMachineRunById.json
func ExampleSoftwareUpdateConfigurationMachineRunsClient_GetByID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSoftwareUpdateConfigurationMachineRunsClient().GetByID(ctx, "mygroup", "myaccount", "ca440719-34a4-4234-a1a9-3f84faf7788f", &armautomation.SoftwareUpdateConfigurationMachineRunsClientGetByIDOptions{ClientRequestID: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SoftwareUpdateConfigurationMachineRun = armautomation.SoftwareUpdateConfigurationMachineRun{
	// 	Name: to.Ptr("ca440719-34a4-4234-a1a9-3f84faf7788f"),
	// 	ID: to.Ptr("/subscriptions/51766542-3ed7-4a72-a187-0c8ab644ddab/resourceGroups/mygroup/providers/Microsoft.Automation/automationAccounts/myaccount/softwareUpdateConfigurationMachineRuns/ca440719-34a4-4234-a1a9-3f84faf7788f"),
	// 	Properties: &armautomation.UpdateConfigurationMachineRunProperties{
	// 		ConfiguredDuration: to.Ptr("PT2H"),
	// 		CorrelationID: to.Ptr("0b943e57-44d3-4f05-898c-6e92aa617e59"),
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-23T02:33:30.748Z"); return t}()),
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-23T02:33:36.416Z"); return t}()),
	// 		Error: &armautomation.ErrorResponse{
	// 		},
	// 		Job: &armautomation.JobNavigation{
	// 		},
	// 		LastModifiedBy: to.Ptr(""),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-23T02:34:32.436Z"); return t}()),
	// 		OSType: to.Ptr("Windows"),
	// 		SoftwareUpdateConfiguration: &armautomation.UpdateConfigurationNavigation{
	// 			Name: to.Ptr("mypatch"),
	// 		},
	// 		SourceComputerID: to.Ptr("3d3f24bf-7037-424e-bfba-aae3b9752f8e"),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-23T02:33:30.748Z"); return t}()),
	// 		Status: to.Ptr("Succeeded"),
	// 		TargetComputer: to.Ptr("/subscriptions/51766542-3ed7-4a72-a187-0c8ab644ddab/resourceGroups/mygroup/providers/Microsoft.Compute/virtualMachines/myvm"),
	// 		TargetComputerType: to.Ptr("AzureVirtualMachines"),
	// 	},
	// }
}
