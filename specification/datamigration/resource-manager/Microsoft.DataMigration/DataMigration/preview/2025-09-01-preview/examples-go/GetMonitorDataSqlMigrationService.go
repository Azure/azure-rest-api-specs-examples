package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v3"
)

// Generated from example definition: 2025-09-01-preview/GetMonitorDataSqlMigrationService.json
func ExampleSQLMigrationServicesClient_ListMonitoringData() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSQLMigrationServicesClient().ListMonitoringData(ctx, "testrg", "service1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatamigration.SQLMigrationServicesClientListMonitoringDataResponse{
	// 	IntegrationRuntimeMonitoringData: armdatamigration.IntegrationRuntimeMonitoringData{
	// 		Name: to.Ptr("IntegrationRuntime1"),
	// 		Nodes: []*armdatamigration.NodeMonitoringData{
	// 			{
	// 				AvailableMemoryInMB: to.Ptr[int32](4219),
	// 				ConcurrentJobsLimit: to.Ptr[int32](20),
	// 				ConcurrentJobsRunning: to.Ptr[int32](0),
	// 				CPUUtilization: to.Ptr[int32](66),
	// 				NodeName: to.Ptr("DESKTOP-6AAAAAA"),
	// 				ReceivedBytes: to.Ptr[float64](0.14946500957012177),
	// 				SentBytes: to.Ptr[float64](0.24564747512340546),
	// 			},
	// 			{
	// 				AvailableMemoryInMB: to.Ptr[int32](4219),
	// 				ConcurrentJobsLimit: to.Ptr[int32](20),
	// 				ConcurrentJobsRunning: to.Ptr[int32](0),
	// 				CPUUtilization: to.Ptr[int32](66),
	// 				NodeName: to.Ptr("DESKTOP-6AAAAAB"),
	// 				ReceivedBytes: to.Ptr[float64](0.14946500957012177),
	// 				SentBytes: to.Ptr[float64](0.24564747512340546),
	// 			},
	// 		},
	// 	},
	// }
}
