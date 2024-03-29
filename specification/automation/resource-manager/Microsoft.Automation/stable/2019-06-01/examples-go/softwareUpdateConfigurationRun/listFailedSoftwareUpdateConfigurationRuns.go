package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfigurationRun/listFailedSoftwareUpdateConfigurationRuns.json
func ExampleSoftwareUpdateConfigurationRunsClient_List_listSoftwareUpdateConfigurationMachineRunWithStatusEqualToFailed() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSoftwareUpdateConfigurationRunsClient().List(ctx, "mygroup", "myaccount", &armautomation.SoftwareUpdateConfigurationRunsClientListOptions{ClientRequestID: nil,
		Filter: to.Ptr("properties/status%20eq%20'Failed'"),
		Skip:   nil,
		Top:    nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SoftwareUpdateConfigurationRunListResult = armautomation.SoftwareUpdateConfigurationRunListResult{
	// 	Value: []*armautomation.SoftwareUpdateConfigurationRun{
	// 		{
	// 			Name: to.Ptr("2bd77cfa-2e9c-41b4-a45b-684a77cfeca9"),
	// 			ID: to.Ptr("/subscriptions/51766542-3ed7-4a72-a187-0c8ab644ddab/resourceGroups/mygroup/providers/Microsoft.Automation/automationAccounts/myaccount/softwareUpdateConfigurationRuns/2bd77cfa-2e9c-41b4-a45b-684a77cfeca9"),
	// 			Properties: &armautomation.SoftwareUpdateConfigurationRunProperties{
	// 				ComputerCount: to.Ptr[int32](1),
	// 				ConfiguredDuration: to.Ptr("PT2H"),
	// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-23T02:30:36.240Z"); return t}()),
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-23T02:30:42.846Z"); return t}()),
	// 				FailedCount: to.Ptr[int32](0),
	// 				LastModifiedBy: to.Ptr(""),
	// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-23T02:31:39.396Z"); return t}()),
	// 				OSType: to.Ptr("Windows"),
	// 				SoftwareUpdateConfiguration: &armautomation.UpdateConfigurationNavigation{
	// 					Name: to.Ptr("mypatch"),
	// 				},
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-23T02:30:36.240Z"); return t}()),
	// 				Status: to.Ptr("Failed"),
	// 				Tasks: &armautomation.SoftwareUpdateConfigurationRunTasks{
	// 					PreTask: &armautomation.SoftwareUpdateConfigurationRunTaskProperties{
	// 						JobID: to.Ptr("be430e9e-2290-462e-8f86-686407c35fab"),
	// 						Source: to.Ptr("preRunbook"),
	// 						Status: to.Ptr("Completed"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("5dabff55-9812-4a58-af16-b0cb1d9384e8"),
	// 			ID: to.Ptr("/subscriptions/51766542-3ed7-4a72-a187-0c8ab644ddab/resourceGroups/mygroup/providers/Microsoft.Automation/automationAccounts/myaccount/softwareUpdateConfigurationRuns/5dabff55-9812-4a58-af16-b0cb1d9384e8"),
	// 			Properties: &armautomation.SoftwareUpdateConfigurationRunProperties{
	// 				ComputerCount: to.Ptr[int32](1),
	// 				ConfiguredDuration: to.Ptr("PT2H"),
	// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-23T01:33:01.881Z"); return t}()),
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-23T01:33:08.113Z"); return t}()),
	// 				FailedCount: to.Ptr[int32](0),
	// 				LastModifiedBy: to.Ptr(""),
	// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-23T01:34:03.940Z"); return t}()),
	// 				OSType: to.Ptr("Windows"),
	// 				SoftwareUpdateConfiguration: &armautomation.UpdateConfigurationNavigation{
	// 					Name: to.Ptr("mypatch"),
	// 				},
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-23T01:33:01.881Z"); return t}()),
	// 				Status: to.Ptr("Failed"),
	// 				Tasks: &armautomation.SoftwareUpdateConfigurationRunTasks{
	// 					PreTask: &armautomation.SoftwareUpdateConfigurationRunTaskProperties{
	// 						JobID: to.Ptr("be430e9e-2290-462e-8f86-686407c35fab"),
	// 						Source: to.Ptr("preRunbook"),
	// 						Status: to.Ptr("Completed"),
	// 					},
	// 				},
	// 			},
	// 	}},
	// }
}
