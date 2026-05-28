package armstoragemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagemover/armstoragemover/v2"
)

// Generated from example definition: 2025-12-01/JobDefinitions_Get_With_Schedule.json
func ExampleJobDefinitionsClient_Get_jobDefinitionsGetWithSchedule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragemover.NewClientFactory("60bcfc77-6589-4da2-b7fd-f9ec9322cf95", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobDefinitionsClient().Get(ctx, "examples-rg", "examples-storageMoverName", "examples-projectName", "examples-jobDefinitionName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstoragemover.JobDefinitionsClientGetResponse{
	// 	JobDefinition: &armstoragemover.JobDefinition{
	// 		Name: to.Ptr("examples-jobDefinitionName"),
	// 		Type: to.Ptr("Microsoft.StorageMover/storageMovers/jobDefinitions"),
	// 		ID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/jobDefinitions/examples-jobDefinitionName"),
	// 		Properties: &armstoragemover.JobDefinitionProperties{
	// 			Description: to.Ptr("Example Job Definition Description"),
	// 			AgentName: to.Ptr("migration-agent"),
	// 			AgentResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/agents/migration-agent"),
	// 			CopyMode: to.Ptr(armstoragemover.CopyModeAdditive),
	// 			SourceName: to.Ptr("examples-sourceEndpointName"),
	// 			SourceResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/examples-sourceEndpointName"),
	// 			SourceSubpath: to.Ptr("/"),
	// 			TargetName: to.Ptr("examples-targetEndpointName"),
	// 			TargetResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/examples-targetEndpointName"),
	// 			TargetSubpath: to.Ptr("/"),
	// 			Connections: []*string{
	// 				to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/connections/example-connection"),
	// 			},
	// 			Schedule: &armstoragemover.ScheduleInfo{
	// 				Frequency: to.Ptr(armstoragemover.FrequencyWeekly),
	// 				IsActive: to.Ptr(true),
	// 				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-12-01T00:00:00Z"); return t}()),
	// 				EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-12-31T12:00:00Z"); return t}()),
	// 				ExecutionTime: &armstoragemover.SchedulerTime{
	// 					Hour: to.Ptr[int32](9),
	// 					Minute: to.Ptr(				armstoragemover.MinuteZero),
	// 				},
	// 				DaysOfWeek: []*string{
	// 					to.Ptr("Monday"),
	// 					to.Ptr("Wednesday"),
	// 					to.Ptr("Friday"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
