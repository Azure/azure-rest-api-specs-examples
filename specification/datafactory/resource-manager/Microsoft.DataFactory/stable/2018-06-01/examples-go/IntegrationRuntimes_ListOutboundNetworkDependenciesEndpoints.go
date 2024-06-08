package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f6f50c6388fd5836fa142384641b8353a99874ef/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_ListOutboundNetworkDependenciesEndpoints.json
func ExampleIntegrationRuntimesClient_ListOutboundNetworkDependenciesEndpoints() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIntegrationRuntimesClient().ListOutboundNetworkDependenciesEndpoints(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleIntegrationRuntime", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IntegrationRuntimeOutboundNetworkDependenciesEndpointsResponse = armdatafactory.IntegrationRuntimeOutboundNetworkDependenciesEndpointsResponse{
	// 	Value: []*armdatafactory.IntegrationRuntimeOutboundNetworkDependenciesCategoryEndpoint{
	// 		{
	// 			Category: to.Ptr("Azure Data Factory (Management)"),
	// 			Endpoints: []*armdatafactory.IntegrationRuntimeOutboundNetworkDependenciesEndpoint{
	// 				{
	// 					DomainName: to.Ptr("wu.frontend.int.clouddatahub-int.net"),
	// 					EndpointDetails: []*armdatafactory.IntegrationRuntimeOutboundNetworkDependenciesEndpointDetails{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 					}},
	// 			}},
	// 		},
	// 		{
	// 			Category: to.Ptr("Azure Storage (Management)"),
	// 			Endpoints: []*armdatafactory.IntegrationRuntimeOutboundNetworkDependenciesEndpoint{
	// 				{
	// 					DomainName: to.Ptr("*.blob.core.windows.net"),
	// 					EndpointDetails: []*armdatafactory.IntegrationRuntimeOutboundNetworkDependenciesEndpointDetails{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("*.table.core.windows.net"),
	// 					EndpointDetails: []*armdatafactory.IntegrationRuntimeOutboundNetworkDependenciesEndpointDetails{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 					}},
	// 			}},
	// 		},
	// 		{
	// 			Category: to.Ptr("Event Hub (Logging)"),
	// 			Endpoints: []*armdatafactory.IntegrationRuntimeOutboundNetworkDependenciesEndpoint{
	// 				{
	// 					DomainName: to.Ptr("*.servicebus.windows.net"),
	// 					EndpointDetails: []*armdatafactory.IntegrationRuntimeOutboundNetworkDependenciesEndpointDetails{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 					}},
	// 			}},
	// 		},
	// 		{
	// 			Category: to.Ptr("Microsoft Logging service (Internal Use)"),
	// 			Endpoints: []*armdatafactory.IntegrationRuntimeOutboundNetworkDependenciesEndpoint{
	// 				{
	// 					DomainName: to.Ptr("gcs.prod.monitoring.core.windows.net"),
	// 					EndpointDetails: []*armdatafactory.IntegrationRuntimeOutboundNetworkDependenciesEndpointDetails{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("prod.warmpath.msftcloudes.com"),
	// 					EndpointDetails: []*armdatafactory.IntegrationRuntimeOutboundNetworkDependenciesEndpointDetails{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("azurewatsonanalysis-prod.core.windows.net"),
	// 					EndpointDetails: []*armdatafactory.IntegrationRuntimeOutboundNetworkDependenciesEndpointDetails{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 					}},
	// 			}},
	// 	}},
	// }
}
