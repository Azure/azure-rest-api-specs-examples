package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfiguration/listSoftwareUpdateConfigurations.json
func ExampleSoftwareUpdateConfigurationsClient_List_listSoftwareUpdateConfigurations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSoftwareUpdateConfigurationsClient().List(ctx, "mygroup", "myaccount", &armautomation.SoftwareUpdateConfigurationsClientListOptions{ClientRequestID: nil,
		Filter: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SoftwareUpdateConfigurationListResult = armautomation.SoftwareUpdateConfigurationListResult{
	// 	Value: []*armautomation.SoftwareUpdateConfigurationCollectionItem{
	// 		{
	// 			Name: to.Ptr("testpatch-01"),
	// 			ID: to.Ptr("/subscriptions/1a7d4044-286c-4acb-969a-96639265bf2e/resourceGroups/Mo-Resources-WCUS/providers/Microsoft.Automation/automationAccounts/Mo-AAA-WCUS/softwareUpdateConfigurations/testpatch-01"),
	// 			Properties: &armautomation.SoftwareUpdateConfigurationCollectionItemProperties{
	// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-19T18:54:50.523Z"); return t}()),
	// 				Frequency: to.Ptr(armautomation.ScheduleFrequencyWeek),
	// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-19T18:54:50.680Z"); return t}()),
	// 				NextRun: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-23T19:22:00.000Z"); return t}()),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-19T19:22:00.000Z"); return t}()),
	// 				Tasks: &armautomation.SoftwareUpdateConfigurationTasks{
	// 					PostTask: &armautomation.TaskProperties{
	// 						Source: to.Ptr("GetCache"),
	// 					},
	// 					PreTask: &armautomation.TaskProperties{
	// 						Parameters: map[string]*string{
	// 							"COMPUTERNAME": to.Ptr("Computer1"),
	// 						},
	// 						Source: to.Ptr("HelloWorld"),
	// 					},
	// 				},
	// 				UpdateConfiguration: &armautomation.UpdateConfiguration{
	// 					AzureVirtualMachines: []*string{
	// 						to.Ptr("/subscriptions/1a7d4044-286c-4acb-969a-96639265bf2e/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-01"),
	// 						to.Ptr("/subscriptions/1a7d4044-286c-4acb-969a-96639265bf2e/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-02"),
	// 						to.Ptr("/subscriptions/1a7d4044-286c-4acb-969a-96639265bf2e/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-03")},
	// 						Duration: to.Ptr("PT2H"),
	// 						OperatingSystem: to.Ptr(armautomation.OperatingSystemTypeWindows),
	// 						Targets: &armautomation.TargetProperties{
	// 							AzureQueries: []*armautomation.AzureQueryProperties{
	// 								{
	// 									Scope: []*string{
	// 										to.Ptr("/subscriptions/422b6c61-95b0-4213-b3be-7282315df71d/resourceGroups/a-stasku-rg0"),
	// 										to.Ptr("/subscriptions/422b6c61-95b0-4213-b3be-7282315df71d")},
	// 										TagSettings: &armautomation.TagSettingsProperties{
	// 											FilterOperator: to.Ptr(armautomation.TagOperatorsAll),
	// 											Tags: map[string][]*string{
	// 												"tag1": []*string{
	// 													to.Ptr("tag1Value1"),
	// 													to.Ptr("tag1Value2")},
	// 													"tag2": []*string{
	// 														to.Ptr("tag2Value1"),
	// 														to.Ptr("tag2Value2")},
	// 													},
	// 												},
	// 										}},
	// 									},
	// 									Windows: &armautomation.WindowsProperties{
	// 										IncludedUpdateClassifications: to.Ptr(armautomation.WindowsUpdateClasses("Critical, Security, UpdateRollup, FeaturePack, ServicePack, Definition, Tools, Updates")),
	// 									},
	// 								},
	// 							},
	// 						},
	// 						{
	// 							Name: to.Ptr("testpatch-02"),
	// 							ID: to.Ptr("/subscriptions/1a7d4044-286c-4acb-969a-96639265bf2e/resourceGroups/Mo-Resources-WCUS/providers/Microsoft.Automation/automationAccounts/Mo-AAA-WCUS/softwareUpdateConfigurations/testpatch-02"),
	// 							Properties: &armautomation.SoftwareUpdateConfigurationCollectionItemProperties{
	// 								CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-11T21:52:02.773Z"); return t}()),
	// 								Frequency: to.Ptr(armautomation.ScheduleFrequencyHour),
	// 								LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-11T21:52:22.880Z"); return t}()),
	// 								NextRun: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-05T19:26:00.000Z"); return t}()),
	// 								ProvisioningState: to.Ptr("Succeeded"),
	// 								StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-05T19:26:00.000Z"); return t}()),
	// 								Tasks: &armautomation.SoftwareUpdateConfigurationTasks{
	// 									PostTask: &armautomation.TaskProperties{
	// 										Source: to.Ptr("GetCache"),
	// 									},
	// 									PreTask: &armautomation.TaskProperties{
	// 										Parameters: map[string]*string{
	// 											"COMPUTERNAME": to.Ptr("Computer1"),
	// 										},
	// 										Source: to.Ptr("HelloWorld"),
	// 									},
	// 								},
	// 								UpdateConfiguration: &armautomation.UpdateConfiguration{
	// 									AzureVirtualMachines: []*string{
	// 										to.Ptr("/subscriptions/1a7d4044-286c-4acb-969a-96639265bf2e/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-04"),
	// 										to.Ptr("/subscriptions/1a7d4044-286c-4acb-969a-96639265bf2e/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-05"),
	// 										to.Ptr("/subscriptions/1a7d4044-286c-4acb-969a-96639265bf2e/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-06")},
	// 										Duration: to.Ptr("PT2H30M"),
	// 										OperatingSystem: to.Ptr(armautomation.OperatingSystemTypeWindows),
	// 										Targets: &armautomation.TargetProperties{
	// 											AzureQueries: []*armautomation.AzureQueryProperties{
	// 												{
	// 													Locations: []*string{
	// 														to.Ptr("Japan East"),
	// 														to.Ptr("UK South")},
	// 														Scope: []*string{
	// 															to.Ptr("/subscriptions/422b6c61-95b0-4213-b3be-7282315df71d/resourceGroups/a-stasku-rg0"),
	// 															to.Ptr("/subscriptions/422b6c61-95b0-4213-b3be-7282315df71d")},
	// 															TagSettings: &armautomation.TagSettingsProperties{
	// 																FilterOperator: to.Ptr(armautomation.TagOperatorsAll),
	// 																Tags: map[string][]*string{
	// 																	"tag1": []*string{
	// 																		to.Ptr("tag1Value1"),
	// 																		to.Ptr("tag1Value2")},
	// 																		"tag2": []*string{
	// 																			to.Ptr("tag2Value1"),
	// 																			to.Ptr("tag2Value2")},
	// 																		},
	// 																	},
	// 															}},
	// 														},
	// 														Windows: &armautomation.WindowsProperties{
	// 															IncludedUpdateClassifications: to.Ptr(armautomation.WindowsUpdateClasses("Critical, FeaturePack")),
	// 														},
	// 													},
	// 												},
	// 										}},
	// 									}
}
