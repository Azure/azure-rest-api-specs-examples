package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2023-05-01-preview/examples/ListOutboundNetworkDependenciesByManagedInstance.json
func ExampleManagedInstancesClient_NewListOutboundNetworkDependenciesByManagedInstancePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedInstancesClient().NewListOutboundNetworkDependenciesByManagedInstancePager("sqlcrudtest-7398", "testinstance", nil)
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
		// page.OutboundEnvironmentEndpointCollection = armsql.OutboundEnvironmentEndpointCollection{
		// 	Value: []*armsql.OutboundEnvironmentEndpoint{
		// 		{
		// 			Category: to.Ptr("Azure SQL Database"),
		// 			Endpoints: []*armsql.EndpointDependency{
		// 				{
		// 					DomainName: to.Ptr("control.database.windows.net"),
		// 					EndpointDetails: []*armsql.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("worker.database.windows.net"),
		// 					EndpointDetails: []*armsql.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 			}},
		// 		},
		// 		{
		// 			Category: to.Ptr("Azure Storage"),
		// 			Endpoints: []*armsql.EndpointDependency{
		// 				{
		// 					DomainName: to.Ptr("blob.core.windows.net"),
		// 					EndpointDetails: []*armsql.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("table.core.windows.net"),
		// 					EndpointDetails: []*armsql.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 			}},
		// 		},
		// 		{
		// 			Category: to.Ptr("Service Bus"),
		// 			Endpoints: []*armsql.EndpointDependency{
		// 				{
		// 					DomainName: to.Ptr("servicebus.windows.net"),
		// 					EndpointDetails: []*armsql.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 			}},
		// 		},
		// 		{
		// 			Category: to.Ptr("Certificate Verification"),
		// 			Endpoints: []*armsql.EndpointDependency{
		// 				{
		// 					DomainName: to.Ptr("dsms.core.windows.net"),
		// 					EndpointDetails: []*armsql.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("dsts.core.windows.net"),
		// 					EndpointDetails: []*armsql.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("login.windows.net"),
		// 					EndpointDetails: []*armsql.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 			}},
		// 		},
		// 		{
		// 			Category: to.Ptr("Telemetry"),
		// 			Endpoints: []*armsql.EndpointDependency{
		// 				{
		// 					DomainName: to.Ptr("azurewatsonanalysis-prod.core.windows.net"),
		// 					EndpointDetails: []*armsql.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("global.metrics.nsatc.net"),
		// 					EndpointDetails: []*armsql.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("production.diagnostics.monitoring.core.windows.net"),
		// 					EndpointDetails: []*armsql.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 			}},
		// 	}},
		// }
	}
}
