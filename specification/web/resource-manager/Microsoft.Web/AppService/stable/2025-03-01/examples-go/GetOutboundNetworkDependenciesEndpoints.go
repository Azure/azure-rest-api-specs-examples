package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/GetOutboundNetworkDependenciesEndpoints.json
func ExampleEnvironmentsClient_NewGetOutboundNetworkDependenciesEndpointsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEnvironmentsClient().NewGetOutboundNetworkDependenciesEndpointsPager("Sample-WestUSResourceGroup", "SampleAse", nil)
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
		// page.OutboundEnvironmentEndpointCollection = armappservice.OutboundEnvironmentEndpointCollection{
		// 	Value: []*armappservice.OutboundEnvironmentEndpoint{
		// 		{
		// 			Category: to.Ptr("Azure Storage"),
		// 			Endpoints: []*armappservice.EndpointDependency{
		// 				{
		// 					DomainName: to.Ptr("blob.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("52.183.104.36"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](42.0469),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("52.183.104.36"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](41.7038),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("queue.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("52.183.104.25"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](37.326),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("52.183.104.25"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](37.513600000000004),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("table.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("52.183.104.26"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](32.789),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("52.183.104.26"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](1.8702),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("file.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("52.183.104.28"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](36.7378),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("52.183.104.28"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](36.7108),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("blob.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.64"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](4.0261000000000005),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.64"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](3.8264),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("queue.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.73"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](40.8523),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.73"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](40.7501),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("table.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.74"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](1.2071),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.74"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](38.2975),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("file.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.76"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](59.383700000000005),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.76"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](60.0775),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("blob.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.64"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](3.5512),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.64"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](2.6777),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("queue.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.73"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](3.5204),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.73"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](2.8193),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("table.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.74"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](3.2720000000000002),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.74"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](3.0147),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("file.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.76"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](3.2387),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.76"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](3.2804),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("blob.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("13.66.176.16"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](43.025200000000005),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("13.66.176.16"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](43.1683),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("queue.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("13.66.176.25"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](41.8598),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("13.66.176.25"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](41.9805),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("table.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("13.66.176.26"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](30.542900000000003),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("13.66.176.26"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](0.9832000000000001),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("file.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("13.66.176.28"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](35.9562),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("13.66.176.28"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](36.0643),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("blob.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.64"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](3.2829),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.64"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](2.3393),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("queue.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.73"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](3.4103000000000003),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.73"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](4.1032),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("table.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.74"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](3.1141),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.74"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](3.0247),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("file.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.76"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](2.834),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.76"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](2.8198000000000003),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("blob.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.64"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](1.3855000000000002),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.64"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](1.0594000000000001),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("queue.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.73"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](2.4717000000000002),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.73"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](2.4827),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("table.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.74"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](3.395),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.74"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](3.1701),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("file.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.76"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](1.4308),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.76"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](22.5866),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("blob.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.64"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](2.5372),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.64"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](2.4626),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("queue.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.73"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](3.7873),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.73"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](4.1911000000000005),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("table.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.74"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](2.9162000000000003),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.74"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](2.7896),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("file.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.76"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](26.723100000000002),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("13.77.184.76"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](26.735200000000003),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 			}},
		// 		},
		// 		{
		// 			Category: to.Ptr("Azure SQL Database"),
		// 			Endpoints: []*armappservice.EndpointDependency{
		// 				{
		// 					DomainName: to.Ptr("database.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("13.66.226.202"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](1.5964),
		// 							Port: to.Ptr[int32](1433),
		// 					}},
		// 			}},
		// 		},
		// 		{
		// 			Category: to.Ptr("Azure Management"),
		// 			Endpoints: []*armappservice.EndpointDependency{
		// 				{
		// 					DomainName: to.Ptr("management.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("23.102.135.246"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](46.5764),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("admin.core.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("23.102.135.247"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](47.408),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("management.azure.com"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("52.151.25.45"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](3.9529),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 			}},
		// 		},
		// 		{
		// 			Category: to.Ptr("Azure Active Directory"),
		// 			Endpoints: []*armappservice.EndpointDependency{
		// 				{
		// 					DomainName: to.Ptr("graph.windows.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("20.190.133.83"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](3.2264),
		// 							Port: to.Ptr[int32](443),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("20.190.133.81"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](3.2264),
		// 							Port: to.Ptr[int32](443),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("20.190.133.67"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](3.2264),
		// 							Port: to.Ptr[int32](443),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("40.126.5.34"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](3.2264),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 			}},
		// 		},
		// 		{
		// 			Category: to.Ptr("Regional Service"),
		// 			Endpoints: []*armappservice.EndpointDependency{
		// 				{
		// 					DomainName: to.Ptr("gr-prod-mwh.cloudapp.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("13.66.225.188"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](3.3826),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("az-prod.metrics.nsatc.net"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("40.77.24.27"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](38.5647),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 			}},
		// 		},
		// 		{
		// 			Category: to.Ptr("SSL Certificate Verification"),
		// 			Endpoints: []*armappservice.EndpointDependency{
		// 				{
		// 					DomainName: to.Ptr("ocsp.msocsp.com"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("104.18.25.243"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](6.0651),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("104.18.25.243"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](12.888),
		// 							Port: to.Ptr[int32](443),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("104.18.24.243"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](6.0651),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("104.18.24.243"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](12.888),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("mscrl.microsoft.com"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("152.199.4.33"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](6.742900000000001),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("152.199.4.33"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](7.436100000000001),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("crl.microsoft.com"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("23.215.102.10"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](25.136200000000002),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("23.215.102.10"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](25.0085),
		// 							Port: to.Ptr[int32](443),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("23.215.102.42"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](25.136200000000002),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("23.215.102.42"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](25.0085),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("www.microsoft.com"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("23.49.13.56"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](7.9229),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("23.49.13.56"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](8.4871),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("crl3.digicert.com"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("72.21.91.29"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](5.4074),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("72.21.91.29"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](5.577),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("ocsp.digicert.com"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("72.21.91.29"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](6.8989),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("72.21.91.29"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](5.667400000000001),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("cacerts.digicert.com"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("104.18.11.39"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](10.772400000000001),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("104.18.11.39"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](10.7705),
		// 							Port: to.Ptr[int32](443),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("104.18.10.39"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](10.772400000000001),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("104.18.10.39"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](10.7705),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("www.thawte.com"),
		// 					EndpointDetails: []*armappservice.EndpointDetail{
		// 						{
		// 							IPAddress: to.Ptr("54.69.98.161"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](47.532900000000005),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("54.69.98.161"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](48.5362),
		// 							Port: to.Ptr[int32](443),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("35.167.62.148"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](47.532900000000005),
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("35.167.62.148"),
		// 							IsAccessible: to.Ptr(true),
		// 							Latency: to.Ptr[float64](48.5362),
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 			}},
		// 	}},
		// }
	}
}
