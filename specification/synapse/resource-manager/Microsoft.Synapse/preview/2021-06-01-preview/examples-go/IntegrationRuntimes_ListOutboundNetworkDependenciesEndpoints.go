package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_ListOutboundNetworkDependenciesEndpoints.json
func ExampleIntegrationRuntimesClient_ListOutboundNetworkDependenciesEndpoints() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewIntegrationRuntimesClient("ade9c2b6-c160-4305-9bb9-80342f6c1ae2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ListOutboundNetworkDependenciesEndpoints(ctx, "exampleResourceGroup", "exampleWorkspace", "exampleIntegrationRuntime", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IntegrationRuntimeOutboundNetworkDependenciesEndpointsResponse = armsynapse.IntegrationRuntimeOutboundNetworkDependenciesEndpointsResponse{
	// 	Value: []*armsynapse.IntegrationRuntimeOutboundNetworkDependenciesCategoryEndpoint{
	// 		{
	// 			Category: to.Ptr("Azure Data Factory (Management)"),
	// 			Endpoints: []*armsynapse.IntegrationRuntimeOutboundNetworkDependenciesEndpoint{
	// 				{
	// 					DomainName: to.Ptr("wu.frontend.int.clouddatahub-int.net"),
	// 					EndpointDetails: []*armsynapse.IntegrationRuntimeOutboundNetworkDependenciesEndpointDetails{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 					}},
	// 			}},
	// 		},
	// 		{
	// 			Category: to.Ptr("Azure Storage (Management)"),
	// 			Endpoints: []*armsynapse.IntegrationRuntimeOutboundNetworkDependenciesEndpoint{
	// 				{
	// 					DomainName: to.Ptr("*.blob.core.windows.net"),
	// 					EndpointDetails: []*armsynapse.IntegrationRuntimeOutboundNetworkDependenciesEndpointDetails{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("*.table.core.windows.net"),
	// 					EndpointDetails: []*armsynapse.IntegrationRuntimeOutboundNetworkDependenciesEndpointDetails{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 					}},
	// 			}},
	// 		},
	// 		{
	// 			Category: to.Ptr("Event Hub (Logging)"),
	// 			Endpoints: []*armsynapse.IntegrationRuntimeOutboundNetworkDependenciesEndpoint{
	// 				{
	// 					DomainName: to.Ptr("*.servicebus.windows.net"),
	// 					EndpointDetails: []*armsynapse.IntegrationRuntimeOutboundNetworkDependenciesEndpointDetails{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 					}},
	// 			}},
	// 		},
	// 		{
	// 			Category: to.Ptr("Microsoft Logging service (Internal Use)"),
	// 			Endpoints: []*armsynapse.IntegrationRuntimeOutboundNetworkDependenciesEndpoint{
	// 				{
	// 					DomainName: to.Ptr("gcs.prod.monitoring.core.windows.net"),
	// 					EndpointDetails: []*armsynapse.IntegrationRuntimeOutboundNetworkDependenciesEndpointDetails{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("prod.warmpath.msftcloudes.com"),
	// 					EndpointDetails: []*armsynapse.IntegrationRuntimeOutboundNetworkDependenciesEndpointDetails{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("azurewatsonanalysis-prod.core.windows.net"),
	// 					EndpointDetails: []*armsynapse.IntegrationRuntimeOutboundNetworkDependenciesEndpointDetails{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 					}},
	// 			}},
	// 	}},
	// }
}
