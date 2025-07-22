package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c548b0bd279f5e233661b1c81fb5b61b19965cd/specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/GetMonitorDataSqlMigrationService.json
func ExampleSQLMigrationServicesClient_ListMonitoringData() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.IntegrationRuntimeMonitoringData = armdatamigration.IntegrationRuntimeMonitoringData{
	// 	Name: to.Ptr("IntegrationRuntime1"),
	// 	Nodes: []*armdatamigration.NodeMonitoringData{
	// 		{
	// 			AvailableMemoryInMB: to.Ptr[int32](4219),
	// 			ConcurrentJobsLimit: to.Ptr[int32](20),
	// 			ConcurrentJobsRunning: to.Ptr[int32](0),
	// 			CPUUtilization: to.Ptr[int32](66),
	// 			NodeName: to.Ptr("DESKTOP-6AAAAAA"),
	// 			ReceivedBytes: to.Ptr[float64](0.14946500957012177),
	// 			SentBytes: to.Ptr[float64](0.24564747512340546),
	// 		},
	// 		{
	// 			AvailableMemoryInMB: to.Ptr[int32](4219),
	// 			ConcurrentJobsLimit: to.Ptr[int32](20),
	// 			ConcurrentJobsRunning: to.Ptr[int32](0),
	// 			CPUUtilization: to.Ptr[int32](66),
	// 			NodeName: to.Ptr("DESKTOP-6AAAAAB"),
	// 			ReceivedBytes: to.Ptr[float64](0.14946500957012177),
	// 			SentBytes: to.Ptr[float64](0.24564747512340546),
	// 	}},
	// }
}
