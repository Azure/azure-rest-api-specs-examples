package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoOutboundNetworkDependenciesList.json
func ExampleClustersClient_NewListOutboundNetworkDependenciesEndpointsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClustersClient().NewListOutboundNetworkDependenciesEndpointsPager("kustorptest", "kustoCluster", nil)
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
		// page.OutboundNetworkDependenciesEndpointListResult = armkusto.OutboundNetworkDependenciesEndpointListResult{
		// 	Value: []*armkusto.OutboundNetworkDependenciesEndpoint{
		// 		{
		// 			Name: to.Ptr("kustoCluster/AzureActiveDirectory"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/OutboundNetworkDependenciesEndpoints"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/OutboundNetworkDependenciesEndpoints/AzureActiveDirectory"),
		// 			Properties: &armkusto.OutboundNetworkDependenciesEndpointProperties{
		// 				Category: to.Ptr("Azure Active Directory"),
		// 				Endpoints: []*armkusto.EndpointDependency{
		// 					{
		// 						DomainName: to.Ptr("login.microsoftonline.com"),
		// 						EndpointDetails: []*armkusto.EndpointDetail{
		// 							{
		// 								IPAddress: to.Ptr("1.2.3.4"),
		// 								Port: to.Ptr[int32](443),
		// 						}},
		// 					},
		// 					{
		// 						DomainName: to.Ptr("graph.windows.net"),
		// 						EndpointDetails: []*armkusto.EndpointDetail{
		// 							{
		// 								IPAddress: to.Ptr("1.2.3.4"),
		// 								Port: to.Ptr[int32](443),
		// 						}},
		// 					},
		// 					{
		// 						DomainName: to.Ptr("graph.microsoft.com"),
		// 						EndpointDetails: []*armkusto.EndpointDetail{
		// 							{
		// 								IPAddress: to.Ptr("1.2.3.4"),
		// 								Port: to.Ptr[int32](443),
		// 						}},
		// 					},
		// 					{
		// 						DomainName: to.Ptr("graph.microsoft-ppe.com"),
		// 						EndpointDetails: []*armkusto.EndpointDetail{
		// 							{
		// 								IPAddress: to.Ptr("1.2.3.4"),
		// 								Port: to.Ptr[int32](443),
		// 						}},
		// 				}},
		// 				ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("kustoCluster/AzureMonitor"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/OutboundNetworkDependenciesEndpoints"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/OutboundNetworkDependenciesEndpoints/AzureMonitor"),
		// 			Properties: &armkusto.OutboundNetworkDependenciesEndpointProperties{
		// 				Category: to.Ptr("Azure Monitor"),
		// 				Endpoints: []*armkusto.EndpointDependency{
		// 					{
		// 						DomainName: to.Ptr("prod.warmpath.msftcloudes.com"),
		// 						EndpointDetails: []*armkusto.EndpointDetail{
		// 							{
		// 								IPAddress: to.Ptr("1.2.3.4"),
		// 								Port: to.Ptr[int32](443),
		// 						}},
		// 					},
		// 					{
		// 						DomainName: to.Ptr("gcs.prod.monitoring.core.windows.net"),
		// 						EndpointDetails: []*armkusto.EndpointDetail{
		// 							{
		// 								IPAddress: to.Ptr("1.2.3.4"),
		// 								Port: to.Ptr[int32](443),
		// 						}},
		// 					},
		// 					{
		// 						DomainName: to.Ptr("production.diagnostics.monitoring.core.windows.net"),
		// 						EndpointDetails: []*armkusto.EndpointDetail{
		// 							{
		// 								IPAddress: to.Ptr("1.2.3.4"),
		// 								Port: to.Ptr[int32](443),
		// 						}},
		// 					},
		// 					{
		// 						DomainName: to.Ptr("shoebox2.metrics.nsatc.net:443"),
		// 						EndpointDetails: []*armkusto.EndpointDetail{
		// 							{
		// 								IPAddress: to.Ptr("1.2.3.4"),
		// 								Port: to.Ptr[int32](443),
		// 						}},
		// 				}},
		// 				ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("kustoCluster/CertificateAuthority"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/OutboundNetworkDependenciesEndpoints"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/OutboundNetworkDependenciesEndpoints/CertificateAuthority"),
		// 			Properties: &armkusto.OutboundNetworkDependenciesEndpointProperties{
		// 				Category: to.Ptr("Certificate Authority"),
		// 				Endpoints: []*armkusto.EndpointDependency{
		// 					{
		// 						DomainName: to.Ptr("ocsp.msocsp.com"),
		// 						EndpointDetails: []*armkusto.EndpointDetail{
		// 							{
		// 								IPAddress: to.Ptr("1.2.3.4"),
		// 								Port: to.Ptr[int32](80),
		// 						}},
		// 					},
		// 					{
		// 						DomainName: to.Ptr("ocsp.digicert.com"),
		// 						EndpointDetails: []*armkusto.EndpointDetail{
		// 							{
		// 								IPAddress: to.Ptr("1.2.3.4"),
		// 								Port: to.Ptr[int32](80),
		// 						}},
		// 					},
		// 					{
		// 						DomainName: to.Ptr("crl3.digicert.com"),
		// 						EndpointDetails: []*armkusto.EndpointDetail{
		// 							{
		// 								IPAddress: to.Ptr("1.2.3.4"),
		// 								Port: to.Ptr[int32](80),
		// 						}},
		// 					},
		// 					{
		// 						DomainName: to.Ptr("crl.microsoft.com"),
		// 						EndpointDetails: []*armkusto.EndpointDetail{
		// 							{
		// 								IPAddress: to.Ptr("1.2.3.4"),
		// 								Port: to.Ptr[int32](80),
		// 						}},
		// 					},
		// 					{
		// 						DomainName: to.Ptr("www.microsoft.com"),
		// 						EndpointDetails: []*armkusto.EndpointDetail{
		// 							{
		// 								IPAddress: to.Ptr("1.2.3.4"),
		// 								Port: to.Ptr[int32](80),
		// 						}},
		// 				}},
		// 				ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("kustoCluster/AzureStorage"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/OutboundNetworkDependenciesEndpoints"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/OutboundNetworkDependenciesEndpoints/AzureStorage"),
		// 			Properties: &armkusto.OutboundNetworkDependenciesEndpointProperties{
		// 				Category: to.Ptr("Azure Storage"),
		// 				Endpoints: []*armkusto.EndpointDependency{
		// 					{
		// 						DomainName: to.Ptr("clusterinternalsa.blob.core.windows.net"),
		// 						EndpointDetails: []*armkusto.EndpointDetail{
		// 							{
		// 								IPAddress: to.Ptr("1.2.3.4"),
		// 								Port: to.Ptr[int32](443),
		// 						}},
		// 					},
		// 					{
		// 						DomainName: to.Ptr("clusterinternalsa.queue.core.windows.net"),
		// 						EndpointDetails: []*armkusto.EndpointDetail{
		// 							{
		// 								IPAddress: to.Ptr("1.2.3.4"),
		// 								Port: to.Ptr[int32](443),
		// 						}},
		// 					},
		// 					{
		// 						DomainName: to.Ptr("clusterinternalsa.table.core.windows.net"),
		// 						EndpointDetails: []*armkusto.EndpointDetail{
		// 							{
		// 								IPAddress: to.Ptr("1.2.3.4"),
		// 								Port: to.Ptr[int32](443),
		// 						}},
		// 				}},
		// 				ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
