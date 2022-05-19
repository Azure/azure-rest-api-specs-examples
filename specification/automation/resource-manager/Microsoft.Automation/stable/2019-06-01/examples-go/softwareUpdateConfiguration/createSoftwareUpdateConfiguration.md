Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fautomation%2Farmautomation%2Fv0.6.0/sdk/resourcemanager/automation/armautomation/README.md) on how to add the SDK to your project and authenticate.

```go
package armautomation_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfiguration/createSoftwareUpdateConfiguration.json
func ExampleSoftwareUpdateConfigurationsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armautomation.NewSoftwareUpdateConfigurationsClient("51766542-3ed7-4a72-a187-0c8ab644ddab", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"mygroup",
		"myaccount",
		"testpatch",
		armautomation.SoftwareUpdateConfiguration{
			Properties: &armautomation.SoftwareUpdateConfigurationProperties{
				ScheduleInfo: &armautomation.SUCScheduleProperties{
					AdvancedSchedule: &armautomation.AdvancedSchedule{
						WeekDays: []*string{
							to.Ptr("Monday"),
							to.Ptr("Thursday")},
					},
					ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-09T11:22:57+00:00"); return t }()),
					Frequency:  to.Ptr(armautomation.ScheduleFrequencyHour),
					Interval:   to.Ptr[int64](1),
					StartTime:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-19T12:22:57+00:00"); return t }()),
					TimeZone:   to.Ptr("America/Los_Angeles"),
				},
				Tasks: &armautomation.SoftwareUpdateConfigurationTasks{
					PostTask: &armautomation.TaskProperties{
						Source: to.Ptr("GetCache"),
					},
					PreTask: &armautomation.TaskProperties{
						Parameters: map[string]*string{
							"COMPUTERNAME": to.Ptr("Computer1"),
						},
						Source: to.Ptr("HelloWorld"),
					},
				},
				UpdateConfiguration: &armautomation.UpdateConfiguration{
					AzureVirtualMachines: []*string{
						to.Ptr("/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-01"),
						to.Ptr("/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-02"),
						to.Ptr("/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-03")},
					Duration: to.Ptr("PT2H0M"),
					NonAzureComputerNames: []*string{
						to.Ptr("box1.contoso.com"),
						to.Ptr("box2.contoso.com")},
					OperatingSystem: to.Ptr(armautomation.OperatingSystemTypeWindows),
					Targets: &armautomation.TargetProperties{
						AzureQueries: []*armautomation.AzureQueryProperties{
							{
								Locations: []*string{
									to.Ptr("Japan East"),
									to.Ptr("UK South")},
								Scope: []*string{
									to.Ptr("/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources"),
									to.Ptr("/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067")},
								TagSettings: &armautomation.TagSettingsProperties{
									FilterOperator: to.Ptr(armautomation.TagOperatorsAll),
									Tags: map[string][]*string{
										"tag1": {
											to.Ptr("tag1Value1"),
											to.Ptr("tag1Value2"),
											to.Ptr("tag1Value3")},
										"tag2": {
											to.Ptr("tag2Value1"),
											to.Ptr("tag2Value2"),
											to.Ptr("tag2Value3")},
									},
								},
							}},
						NonAzureQueries: []*armautomation.NonAzureQueryProperties{
							{
								FunctionAlias: to.Ptr("SavedSearch1"),
								WorkspaceID:   to.Ptr("WorkspaceId1"),
							},
							{
								FunctionAlias: to.Ptr("SavedSearch2"),
								WorkspaceID:   to.Ptr("WorkspaceId2"),
							}},
					},
					Windows: &armautomation.WindowsProperties{
						ExcludedKbNumbers: []*string{
							to.Ptr("168934"),
							to.Ptr("168973")},
						IncludedUpdateClassifications: to.Ptr(armautomation.WindowsUpdateClassesCritical),
						RebootSetting:                 to.Ptr("IfRequired"),
					},
				},
			},
		},
		&armautomation.SoftwareUpdateConfigurationsClientCreateOptions{ClientRequestID: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
