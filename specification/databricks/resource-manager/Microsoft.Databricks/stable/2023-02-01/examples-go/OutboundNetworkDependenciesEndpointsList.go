package armdatabricks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databricks/armdatabricks"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e1a87e1a5deb3f986ea1474d233d6680f1e19fc1/specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/OutboundNetworkDependenciesEndpointsList.json
func ExampleOutboundNetworkDependenciesEndpointsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabricks.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOutboundNetworkDependenciesEndpointsClient().List(ctx, "myResourceGroup", "myWorkspace", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OutboundEnvironmentEndpointArray = []*armdatabricks.OutboundEnvironmentEndpoint{
	// 	{
	// 		Category: to.Ptr("Webapp"),
	// 		Endpoints: []*armdatabricks.EndpointDependency{
	// 			{
	// 				EndpointDetails: []*armdatabricks.EndpointDetail{
	// 					{
	// 						IPAddress: to.Ptr("11.111.111.11/11"),
	// 						Port: to.Ptr[int32](123),
	// 					},
	// 					{
	// 						IPAddress: to.Ptr("22.222.222.22/22"),
	// 						Port: to.Ptr[int32](123),
	// 				}},
	// 		}},
	// 	},
	// 	{
	// 		Category: to.Ptr("Control Plane NAT"),
	// 		Endpoints: []*armdatabricks.EndpointDependency{
	// 			{
	// 				EndpointDetails: []*armdatabricks.EndpointDetail{
	// 					{
	// 						IPAddress: to.Ptr("33.33.333.333/33"),
	// 						Port: to.Ptr[int32](123),
	// 				}},
	// 		}},
	// 	},
	// 	{
	// 		Category: to.Ptr("Extended infrastructure"),
	// 		Endpoints: []*armdatabricks.EndpointDependency{
	// 			{
	// 				EndpointDetails: []*armdatabricks.EndpointDetail{
	// 					{
	// 						IPAddress: to.Ptr("44.44.44.44/44"),
	// 						Port: to.Ptr[int32](123),
	// 				}},
	// 		}},
	// 	},
	// 	{
	// 		Category: to.Ptr("Azure Storage"),
	// 		Endpoints: []*armdatabricks.EndpointDependency{
	// 			{
	// 				DomainName: to.Ptr("xxx.blob.core.windows.net"),
	// 				EndpointDetails: []*armdatabricks.EndpointDetail{
	// 					{
	// 						Port: to.Ptr[int32](123),
	// 				}},
	// 			},
	// 			{
	// 				DomainName: to.Ptr("yyy.blob.core.windows.net"),
	// 				EndpointDetails: []*armdatabricks.EndpointDetail{
	// 					{
	// 						Port: to.Ptr[int32](123),
	// 				}},
	// 			},
	// 			{
	// 				DomainName: to.Ptr("zzz.blob.core.windows.net"),
	// 				EndpointDetails: []*armdatabricks.EndpointDetail{
	// 					{
	// 						Port: to.Ptr[int32](123),
	// 				}},
	// 		}},
	// 	},
	// 	{
	// 		Category: to.Ptr("Azure My SQL"),
	// 		Endpoints: []*armdatabricks.EndpointDependency{
	// 			{
	// 				DomainName: to.Ptr("xxx.mysql.database.azure.com"),
	// 				EndpointDetails: []*armdatabricks.EndpointDetail{
	// 					{
	// 						Port: to.Ptr[int32](1234),
	// 				}},
	// 			},
	// 			{
	// 				DomainName: to.Ptr("yyy.mysql.database.azure.com"),
	// 				EndpointDetails: []*armdatabricks.EndpointDetail{
	// 					{
	// 						Port: to.Ptr[int32](1234),
	// 				}},
	// 		}},
	// 	},
	// 	{
	// 		Category: to.Ptr("Azure Servicebus"),
	// 		Endpoints: []*armdatabricks.EndpointDependency{
	// 			{
	// 				DomainName: to.Ptr("xxx.servicebus.windows.net"),
	// 				EndpointDetails: []*armdatabricks.EndpointDetail{
	// 					{
	// 						Port: to.Ptr[int32](1234),
	// 				}},
	// 		}},
	// }}
}
