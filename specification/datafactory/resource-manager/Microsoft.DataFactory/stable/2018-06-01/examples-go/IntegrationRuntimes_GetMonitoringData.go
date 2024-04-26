package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_GetMonitoringData.json
func ExampleIntegrationRuntimesClient_GetMonitoringData() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIntegrationRuntimesClient().GetMonitoringData(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleIntegrationRuntime", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IntegrationRuntimeMonitoringData = armdatafactory.IntegrationRuntimeMonitoringData{
	// 	Name: to.Ptr("exampleIntegrationRuntime"),
	// 	Nodes: []*armdatafactory.IntegrationRuntimeNodeMonitoringData{
	// 		{
	// 			AvailableMemoryInMB: to.Ptr[int32](16740),
	// 			ConcurrentJobsLimit: to.Ptr[int32](28),
	// 			ConcurrentJobsRunning: to.Ptr[int32](0),
	// 			CPUUtilization: to.Ptr[int32](15),
	// 			NodeName: to.Ptr("Node_1"),
	// 			ReceivedBytes: to.Ptr[float32](6.731423377990723),
	// 			SentBytes: to.Ptr[float32](2.647491693496704),
	// 	}},
	// }
}
