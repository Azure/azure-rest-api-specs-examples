package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfiguration/listSoftwareUpdateConfigurationsByVm.json
func ExampleSoftwareUpdateConfigurationsClient_List_listSoftwareUpdateConfigurationsTargetingASpecificAzureVirtualMachine() {
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
		Filter: to.Ptr("properties/updateConfiguration/azureVirtualMachines/any(m: m eq '/subscriptions/1a7d4044-286c-4acb-969a-96639265bf2e/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-01')"),
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
	// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-19T18:54:50.5233333+00:00"); return t}()),
	// 				Frequency: to.Ptr(armautomation.ScheduleFrequencyWeek),
	// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-19T18:54:50.68+00:00"); return t}()),
	// 				NextRun: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-23T12:22:00-07:00"); return t}()),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-19T12:22:00-07:00"); return t}()),
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
	// 						Windows: &armautomation.WindowsProperties{
	// 							IncludedUpdateClassifications: to.Ptr(armautomation.WindowsUpdateClasses("Critical, Security, UpdateRollup, FeaturePack, ServicePack, Definition, Tools, Updates")),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			{
	// 				Name: to.Ptr("testpatch-02"),
	// 				ID: to.Ptr("/subscriptions/1a7d4044-286c-4acb-969a-96639265bf2e/resourceGroups/Mo-Resources-WCUS/providers/Microsoft.Automation/automationAccounts/Mo-AAA-WCUS/softwareUpdateConfigurations/testpatch-02"),
	// 				Properties: &armautomation.SoftwareUpdateConfigurationCollectionItemProperties{
	// 					CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-11T21:52:02.7733333+00:00"); return t}()),
	// 					Frequency: to.Ptr(armautomation.ScheduleFrequencyHour),
	// 					LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-11T21:52:22.88+00:00"); return t}()),
	// 					NextRun: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-05T12:26:00-07:00"); return t}()),
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-05T12:26:00-07:00"); return t}()),
	// 					Tasks: &armautomation.SoftwareUpdateConfigurationTasks{
	// 						PostTask: &armautomation.TaskProperties{
	// 							Source: to.Ptr("GetCache"),
	// 						},
	// 						PreTask: &armautomation.TaskProperties{
	// 							Parameters: map[string]*string{
	// 								"COMPUTERNAME": to.Ptr("Computer1"),
	// 							},
	// 							Source: to.Ptr("HelloWorld"),
	// 						},
	// 					},
	// 					UpdateConfiguration: &armautomation.UpdateConfiguration{
	// 						AzureVirtualMachines: []*string{
	// 							to.Ptr("/subscriptions/1a7d4044-286c-4acb-969a-96639265bf2e/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-01"),
	// 							to.Ptr("/subscriptions/1a7d4044-286c-4acb-969a-96639265bf2e/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-05"),
	// 							to.Ptr("/subscriptions/1a7d4044-286c-4acb-969a-96639265bf2e/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-06")},
	// 							Duration: to.Ptr("PT2H30M"),
	// 							OperatingSystem: to.Ptr(armautomation.OperatingSystemTypeWindows),
	// 							Windows: &armautomation.WindowsProperties{
	// 								IncludedUpdateClassifications: to.Ptr(armautomation.WindowsUpdateClasses("Critical, FeaturePack")),
	// 							},
	// 						},
	// 					},
	// 			}},
	// 		}
}
