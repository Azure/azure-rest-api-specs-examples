package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfiguration/getSoftwareUpdateConfigurationByName.json
func ExampleSoftwareUpdateConfigurationsClient_GetByName() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSoftwareUpdateConfigurationsClient().GetByName(ctx, "mygroup", "myaccount", "mypatch", &armautomation.SoftwareUpdateConfigurationsClientGetByNameOptions{ClientRequestID: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SoftwareUpdateConfiguration = armautomation.SoftwareUpdateConfiguration{
	// 	Name: to.Ptr("testpatch"),
	// 	ID: to.Ptr("/subscriptions/51766542-3ed7-4a72-a187-0c8ab644ddab/resourceGroups/mygroup/providers/Microsoft.Automation/automationAccounts/myaccount/softwareUpdateConfigurations/testpatch"),
	// 	Properties: &armautomation.SoftwareUpdateConfigurationProperties{
	// 		CreatedBy: to.Ptr("eve@contoso.com"),
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-19T18:54:50.523Z"); return t}()),
	// 		Error: &armautomation.ErrorResponse{
	// 		},
	// 		LastModifiedBy: to.Ptr(""),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-19T18:54:50.680Z"); return t}()),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ScheduleInfo: &armautomation.SUCScheduleProperties{
	// 			Description: to.Ptr(""),
	// 			CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-19T18:54:50.523Z"); return t}()),
	// 			ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-09T19:22:00.000Z"); return t}()),
	// 			ExpiryTimeOffsetMinutes: to.Ptr[float64](-480),
	// 			Frequency: to.Ptr(armautomation.ScheduleFrequencyWeek),
	// 			Interval: to.Ptr[int64](1),
	// 			IsEnabled: to.Ptr(true),
	// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-19T18:54:50.523Z"); return t}()),
	// 			NextRun: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-19T19:22:00.000Z"); return t}()),
	// 			NextRunOffsetMinutes: to.Ptr[float64](-420),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-19T19:22:00.000Z"); return t}()),
	// 			StartTimeOffsetMinutes: to.Ptr[float64](-420),
	// 			TimeZone: to.Ptr("America/Los_Angeles"),
	// 		},
	// 		Tasks: &armautomation.SoftwareUpdateConfigurationTasks{
	// 			PostTask: &armautomation.TaskProperties{
	// 				Source: to.Ptr("GetCache"),
	// 			},
	// 			PreTask: &armautomation.TaskProperties{
	// 				Parameters: map[string]*string{
	// 					"COMPUTERNAME": to.Ptr("Computer1"),
	// 				},
	// 				Source: to.Ptr("HelloWorld"),
	// 			},
	// 		},
	// 		UpdateConfiguration: &armautomation.UpdateConfiguration{
	// 			AzureVirtualMachines: []*string{
	// 				to.Ptr("/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-01"),
	// 				to.Ptr("/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-02"),
	// 				to.Ptr("/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-03")},
	// 				Duration: to.Ptr("PT2H"),
	// 				Linux: &armautomation.LinuxProperties{
	// 				},
	// 				NonAzureComputerNames: []*string{
	// 					to.Ptr("box1.contoso.com"),
	// 					to.Ptr("box2.contoso.com")},
	// 					OperatingSystem: to.Ptr(armautomation.OperatingSystemTypeWindows),
	// 					Targets: &armautomation.TargetProperties{
	// 						AzureQueries: []*armautomation.AzureQueryProperties{
	// 							{
	// 								Scope: []*string{
	// 									to.Ptr("/subscriptions/422b6c61-95b0-4213-b3be-7282315df71d/resourceGroups/a-stasku-rg0"),
	// 									to.Ptr("/subscriptions/422b6c61-95b0-4213-b3be-7282315df71d")},
	// 									TagSettings: &armautomation.TagSettingsProperties{
	// 										FilterOperator: to.Ptr(armautomation.TagOperatorsAll),
	// 										Tags: map[string][]*string{
	// 											"tag1": []*string{
	// 												to.Ptr("tag1Value1"),
	// 												to.Ptr("tag1Value2")},
	// 												"tag2": []*string{
	// 													to.Ptr("tag2Value1"),
	// 													to.Ptr("tag2Value2")},
	// 												},
	// 											},
	// 									}},
	// 								},
	// 								Windows: &armautomation.WindowsProperties{
	// 									ExcludedKbNumbers: []*string{
	// 										to.Ptr("168934"),
	// 										to.Ptr("168973")},
	// 										IncludedUpdateClassifications: to.Ptr(armautomation.WindowsUpdateClassesCritical),
	// 									},
	// 								},
	// 							},
	// 						}
}
