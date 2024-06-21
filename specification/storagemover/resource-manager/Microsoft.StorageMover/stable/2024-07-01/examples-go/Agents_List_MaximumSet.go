package armstoragemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagemover/armstoragemover/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4ee6d9fd7687d4b67117c5a167c191a7e7e70b53/specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2024-07-01/examples/Agents_List_MaximumSet.json
func ExampleAgentsClient_NewListPager_agentsListMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAgentsClient().NewListPager("examples-rg", "examples-storageMoverName", nil)
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
		// page.AgentList = armstoragemover.AgentList{
		// 	Value: []*armstoragemover.Agent{
		// 		{
		// 			Name: to.Ptr("examples-agentName1"),
		// 			Type: to.Ptr("Microsoft.StorageMover/storageMovers/agents"),
		// 			ID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/agents/examples-agentName1"),
		// 			SystemData: &armstoragemover.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T01:01:01.107Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armstoragemover.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T02:01:01.107Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armstoragemover.CreatedByTypeUser),
		// 			},
		// 			Properties: &armstoragemover.AgentProperties{
		// 				Description: to.Ptr("Example Agent 1 Description"),
		// 				AgentStatus: to.Ptr(armstoragemover.AgentStatusOnline),
		// 				AgentVersion: to.Ptr("1.0.0"),
		// 				ArcResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.HybridCompute/machines/examples-hybridComputeName1"),
		// 				ArcVMUUID: to.Ptr("3bb2c024-eba9-4d18-9e7a-1d772fcc5fe9"),
		// 				LastStatusUpdate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T02:21:01.107Z"); return t}()),
		// 				LocalIPAddress: to.Ptr("192.168.0.0"),
		// 				MemoryInMB: to.Ptr[int64](4096),
		// 				NumberOfCores: to.Ptr[int64](8),
		// 				ProvisioningState: to.Ptr(armstoragemover.ProvisioningStateSucceeded),
		// 				TimeZone: to.Ptr("Eastern Standard Time"),
		// 				UploadLimitSchedule: &armstoragemover.UploadLimitSchedule{
		// 					WeeklyRecurrences: []*armstoragemover.UploadLimitWeeklyRecurrence{
		// 						{
		// 							LimitInMbps: to.Ptr[int32](2000),
		// 							EndTime: &armstoragemover.Time{
		// 								Hour: to.Ptr[int32](18),
		// 								Minute: to.Ptr(armstoragemover.Minute(30)),
		// 							},
		// 							StartTime: &armstoragemover.Time{
		// 								Hour: to.Ptr[int32](9),
		// 								Minute: to.Ptr(armstoragemover.Minute(0)),
		// 							},
		// 							Days: []*armstoragemover.DayOfWeek{
		// 								to.Ptr(armstoragemover.DayOfWeekMonday)},
		// 						}},
		// 					},
		// 					UptimeInSeconds: to.Ptr[int64](522),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("examples-agentName2"),
		// 				Type: to.Ptr("Microsoft.StorageMover/storageMovers/agents"),
		// 				ID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/agents/examples-agentName2"),
		// 				SystemData: &armstoragemover.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T01:01:01.107Z"); return t}()),
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armstoragemover.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T02:01:01.107Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armstoragemover.CreatedByTypeUser),
		// 				},
		// 				Properties: &armstoragemover.AgentProperties{
		// 					AgentStatus: to.Ptr(armstoragemover.AgentStatusOnline),
		// 					AgentVersion: to.Ptr("1.0.0"),
		// 					ArcResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.HybridCompute/machines/examples-hybridComputeName2"),
		// 					ArcVMUUID: to.Ptr("147a1f84-7bbf-4e99-9a6a-a1735a91dfd5"),
		// 					LastStatusUpdate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T02:21:01.107Z"); return t}()),
		// 					LocalIPAddress: to.Ptr("192.168.0.0"),
		// 					MemoryInMB: to.Ptr[int64](4096),
		// 					NumberOfCores: to.Ptr[int64](8),
		// 					ProvisioningState: to.Ptr(armstoragemover.ProvisioningStateSucceeded),
		// 					TimeZone: to.Ptr("Eastern Standard Time"),
		// 					UptimeInSeconds: to.Ptr[int64](877),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("examples-agentName3"),
		// 				Type: to.Ptr("Microsoft.StorageMover/storageMovers/agents"),
		// 				ID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/agents/examples-agentName3"),
		// 				SystemData: &armstoragemover.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T01:01:01.107Z"); return t}()),
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armstoragemover.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T02:01:01.107Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armstoragemover.CreatedByTypeUser),
		// 				},
		// 				Properties: &armstoragemover.AgentProperties{
		// 					AgentStatus: to.Ptr(armstoragemover.AgentStatusOnline),
		// 					AgentVersion: to.Ptr("1.0.0"),
		// 					ArcResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.HybridCompute/machines/examples-hybridComputeName3"),
		// 					ArcVMUUID: to.Ptr("648a7958-f99e-4268-b20e-94c96558dc0d"),
		// 					ErrorDetails: &armstoragemover.AgentPropertiesErrorDetails{
		// 						Code: to.Ptr("SampleErrorCode"),
		// 						Message: to.Ptr("Detailed sample error message."),
		// 					},
		// 					LastStatusUpdate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T02:21:01.107Z"); return t}()),
		// 					LocalIPAddress: to.Ptr("192.168.0.0"),
		// 					MemoryInMB: to.Ptr[int64](100024),
		// 					NumberOfCores: to.Ptr[int64](32),
		// 					ProvisioningState: to.Ptr(armstoragemover.ProvisioningStateSucceeded),
		// 					TimeZone: to.Ptr("Eastern Standard Time"),
		// 					UploadLimitSchedule: &armstoragemover.UploadLimitSchedule{
		// 						WeeklyRecurrences: []*armstoragemover.UploadLimitWeeklyRecurrence{
		// 							{
		// 								LimitInMbps: to.Ptr[int32](5000),
		// 								EndTime: &armstoragemover.Time{
		// 									Hour: to.Ptr[int32](24),
		// 									Minute: to.Ptr(armstoragemover.Minute(0)),
		// 								},
		// 								StartTime: &armstoragemover.Time{
		// 									Hour: to.Ptr[int32](0),
		// 									Minute: to.Ptr(armstoragemover.Minute(0)),
		// 								},
		// 								Days: []*armstoragemover.DayOfWeek{
		// 									to.Ptr(armstoragemover.DayOfWeekSaturday),
		// 									to.Ptr(armstoragemover.DayOfWeekSunday)},
		// 							}},
		// 						},
		// 						UptimeInSeconds: to.Ptr[int64](1025),
		// 					},
		// 			}},
		// 		}
	}
}
