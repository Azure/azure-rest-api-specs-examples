package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementServiceGetOutboundNetworkDependenciesEndpoints.json
func ExampleOutboundNetworkDependenciesEndpointsClient_ListByService() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOutboundNetworkDependenciesEndpointsClient().ListByService(ctx, "rg1", "apimService1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OutboundEnvironmentEndpointList = armapimanagement.OutboundEnvironmentEndpointList{
	// 	Value: []*armapimanagement.OutboundEnvironmentEndpoint{
	// 		{
	// 			Category: to.Ptr("Azure SMTP"),
	// 			Endpoints: []*armapimanagement.EndpointDependency{
	// 				{
	// 					DomainName: to.Ptr("smtpi-ch1.msn.com"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](25028),
	// 							Region: to.Ptr("West US"),
	// 					}},
	// 			}},
	// 		},
	// 		{
	// 			Category: to.Ptr("Azure SQL"),
	// 			Endpoints: []*armapimanagement.EndpointDependency{
	// 				{
	// 					DomainName: to.Ptr("xxxx1345234.database.windows.net"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](1433),
	// 							Region: to.Ptr("West US"),
	// 					}},
	// 			}},
	// 		},
	// 		{
	// 			Category: to.Ptr("Azure Storage"),
	// 			Endpoints: []*armapimanagement.EndpointDependency{
	// 				{
	// 					DomainName: to.Ptr("xxxx32storagedgfbay.blob.core.windows.net"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("West US"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("xxxx1362629927xt.blob.core.windows.net"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("West US"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("xxxx1362629927xt.table.core.windows.net"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("West US"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("xxxx141483183xt.blob.core.windows.net"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("West US"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("xxxx141483183xt.table.core.windows.net"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("West US"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("xxxx1949864718xt.blob.core.windows.net"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("West US"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("xxxx1949864718xt.table.core.windows.net"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("West US"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("xxxx3292114122xt.blob.core.windows.net"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("West US"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("xxxx3292114122xt.table.core.windows.net"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("West US"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("xxxx32tst4oto8t0mlesawmm.blob.core.windows.net"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("West US"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("xxxx32tst4oto8t0mlesawmm.file.core.windows.net"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](445),
	// 							Region: to.Ptr("West US"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("xxxx32tst4oto8t0mlesawmm.queue.core.windows.net"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("West US"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("xxxx32tst4oto8t0mlesawmm.table.core.windows.net"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("West US"),
	// 					}},
	// 			}},
	// 		},
	// 		{
	// 			Category: to.Ptr("Azure Event Hub"),
	// 			Endpoints: []*armapimanagement.EndpointDependency{
	// 				{
	// 					DomainName: to.Ptr("xxxx1362629927eh.servicebus.windows.net"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("West US"),
	// 						},
	// 						{
	// 							Port: to.Ptr[int32](5671),
	// 							Region: to.Ptr("West US"),
	// 						},
	// 						{
	// 							Port: to.Ptr[int32](5672),
	// 							Region: to.Ptr("West US"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("xxxx1949864718eh.servicebus.windows.net"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("West US"),
	// 						},
	// 						{
	// 							Port: to.Ptr[int32](5671),
	// 							Region: to.Ptr("West US"),
	// 						},
	// 						{
	// 							Port: to.Ptr[int32](5672),
	// 							Region: to.Ptr("West US"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("xxxx3292114122eh.servicebus.windows.net"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("West US"),
	// 						},
	// 						{
	// 							Port: to.Ptr[int32](5671),
	// 							Region: to.Ptr("West US"),
	// 						},
	// 						{
	// 							Port: to.Ptr[int32](5672),
	// 							Region: to.Ptr("West US"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("xxxx141483183eh.servicebus.windows.net"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("West US"),
	// 						},
	// 						{
	// 							Port: to.Ptr[int32](5671),
	// 							Region: to.Ptr("West US"),
	// 						},
	// 						{
	// 							Port: to.Ptr[int32](5672),
	// 							Region: to.Ptr("West US"),
	// 					}},
	// 			}},
	// 		},
	// 		{
	// 			Category: to.Ptr("SSL Certificate Verification"),
	// 			Endpoints: []*armapimanagement.EndpointDependency{
	// 				{
	// 					DomainName: to.Ptr("ocsp.msocsp.com"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](80),
	// 							Region: to.Ptr("Global"),
	// 						},
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("Global"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("mscrl.microsoft.com"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](80),
	// 							Region: to.Ptr("Global"),
	// 						},
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("Global"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("crl.microsoft.com"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](80),
	// 							Region: to.Ptr("Global"),
	// 						},
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("Global"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("crl3.digicert.com"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](80),
	// 							Region: to.Ptr("Global"),
	// 						},
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("Global"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("ocsp.digicert.com"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](80),
	// 							Region: to.Ptr("Global"),
	// 						},
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("Global"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("cacerts.digicert.com"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](80),
	// 							Region: to.Ptr("Global"),
	// 						},
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("Global"),
	// 					}},
	// 			}},
	// 		},
	// 		{
	// 			Category: to.Ptr("Azure Monitor"),
	// 			Endpoints: []*armapimanagement.EndpointDependency{
	// 				{
	// 					DomainName: to.Ptr("gcs.ppe.monitoring.core.windows.net"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("Global"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("global.prod.microsoftmetrics.com"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("Global"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("xxx3.prod.microsoftmetrics.com"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](1886),
	// 							Region: to.Ptr("Global"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("xxx3-red.prod.microsoftmetrics.com"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](1886),
	// 							Region: to.Ptr("Global"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("xxx3-black.prod.microsoftmetrics.com"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](1886),
	// 							Region: to.Ptr("Global"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("gcs.ppe.warm.ingestion.monitoring.azure.com"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("Global"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("metrichost23.prod.microsoftmetrics.com"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("Global"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("metrichost23-red.prod.microsoftmetrics.com"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("Global"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("metrichost23-black.prod.microsoftmetrics.com"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("Global"),
	// 					}},
	// 			}},
	// 		},
	// 		{
	// 			Category: to.Ptr("Portal Captcha"),
	// 			Endpoints: []*armapimanagement.EndpointDependency{
	// 				{
	// 					DomainName: to.Ptr("client.xxx.live.com"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("Global"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("partner.xxx.live.com"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("Global"),
	// 					}},
	// 			}},
	// 		},
	// 		{
	// 			Category: to.Ptr("Azure Active Directory"),
	// 			Endpoints: []*armapimanagement.EndpointDependency{
	// 				{
	// 					DomainName: to.Ptr("login.windows.net"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("Global"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("graph.windows.net"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("Global"),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("login.microsoftonline.com"),
	// 					EndpointDetails: []*armapimanagement.EndpointDetail{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 							Region: to.Ptr("Global"),
	// 					}},
	// 			}},
	// 	}},
	// }
}
