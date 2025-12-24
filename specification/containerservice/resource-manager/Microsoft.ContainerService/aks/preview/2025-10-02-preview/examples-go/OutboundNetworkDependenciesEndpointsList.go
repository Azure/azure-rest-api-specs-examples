package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/455d20a5e76d8184f7cff960501a57e1f88986b7/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2025-10-02-preview/examples/OutboundNetworkDependenciesEndpointsList.json
func ExampleManagedClustersClient_NewListOutboundNetworkDependenciesEndpointsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedClustersClient().NewListOutboundNetworkDependenciesEndpointsPager("rg1", "clustername1", nil)
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
		// page.OutboundEnvironmentEndpointCollection = armcontainerservice.OutboundEnvironmentEndpointCollection{
		// 	Value: []*armcontainerservice.OutboundEnvironmentEndpoint{
		// 		{
		// 			Category: to.Ptr("azure-resource-management"),
		// 			Endpoints: []*armcontainerservice.EndpointDependency{
		// 				{
		// 					DomainName: to.Ptr("management.azure.com"),
		// 					EndpointDetails: []*armcontainerservice.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 							Protocol: to.Ptr("Https"),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("login.microsoftonline.com"),
		// 					EndpointDetails: []*armcontainerservice.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 							Protocol: to.Ptr("Https"),
		// 					}},
		// 			}},
		// 		},
		// 		{
		// 			Category: to.Ptr("images"),
		// 			Endpoints: []*armcontainerservice.EndpointDependency{
		// 				{
		// 					DomainName: to.Ptr("mcr.microsoft.com"),
		// 					EndpointDetails: []*armcontainerservice.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 							Protocol: to.Ptr("Https"),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("*.data.mcr.microsoft.com"),
		// 					EndpointDetails: []*armcontainerservice.EndpointDetail{
		// 						{
		// 							Description: to.Ptr("mcr cdn"),
		// 							Port: to.Ptr[int32](443),
		// 							Protocol: to.Ptr("Https"),
		// 					}},
		// 			}},
		// 		},
		// 		{
		// 			Category: to.Ptr("artifacts"),
		// 			Endpoints: []*armcontainerservice.EndpointDependency{
		// 				{
		// 					DomainName: to.Ptr("packages.microsoft.com"),
		// 					EndpointDetails: []*armcontainerservice.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 							Protocol: to.Ptr("Https"),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("acs-mirror.azureedge.net"),
		// 					EndpointDetails: []*armcontainerservice.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 							Protocol: to.Ptr("Https"),
		// 					}},
		// 			}},
		// 		},
		// 		{
		// 			Category: to.Ptr("time-sync"),
		// 			Endpoints: []*armcontainerservice.EndpointDependency{
		// 				{
		// 					DomainName: to.Ptr("ntp.ubuntu.com"),
		// 					EndpointDetails: []*armcontainerservice.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](123),
		// 							Protocol: to.Ptr("UDP"),
		// 					}},
		// 			}},
		// 		},
		// 		{
		// 			Category: to.Ptr("ubuntu-optional"),
		// 			Endpoints: []*armcontainerservice.EndpointDependency{
		// 				{
		// 					DomainName: to.Ptr("security.ubuntu.com"),
		// 					EndpointDetails: []*armcontainerservice.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](80),
		// 							Protocol: to.Ptr("Http"),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("azure.archive.ubuntu.com"),
		// 					EndpointDetails: []*armcontainerservice.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](80),
		// 							Protocol: to.Ptr("Http"),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("changelogs.ubuntu.com"),
		// 					EndpointDetails: []*armcontainerservice.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](80),
		// 							Protocol: to.Ptr("Http"),
		// 					}},
		// 			}},
		// 		},
		// 		{
		// 			Category: to.Ptr("gpu"),
		// 			Endpoints: []*armcontainerservice.EndpointDependency{
		// 				{
		// 					DomainName: to.Ptr("nvidia.github.io"),
		// 					EndpointDetails: []*armcontainerservice.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 							Protocol: to.Ptr("Https"),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("us.download.nvidia.com"),
		// 					EndpointDetails: []*armcontainerservice.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 							Protocol: to.Ptr("Https"),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("apt.dockerproject.org"),
		// 					EndpointDetails: []*armcontainerservice.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 							Protocol: to.Ptr("Https"),
		// 					}},
		// 			}},
		// 		},
		// 		{
		// 			Category: to.Ptr("windows"),
		// 			Endpoints: []*armcontainerservice.EndpointDependency{
		// 				{
		// 					DomainName: to.Ptr("onegetcdn.azureedge.net"),
		// 					EndpointDetails: []*armcontainerservice.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 							Protocol: to.Ptr("Https"),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("go.microsoft.com"),
		// 					EndpointDetails: []*armcontainerservice.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 							Protocol: to.Ptr("Https"),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("*.mp.microsoft.com"),
		// 					EndpointDetails: []*armcontainerservice.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](80),
		// 							Protocol: to.Ptr("Http"),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("www.msftconnecttest.com"),
		// 					EndpointDetails: []*armcontainerservice.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](80),
		// 							Protocol: to.Ptr("Http"),
		// 					}},
		// 				},
		// 				{
		// 					DomainName: to.Ptr("ctldl.windowsupdate.com"),
		// 					EndpointDetails: []*armcontainerservice.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](80),
		// 							Protocol: to.Ptr("Http"),
		// 					}},
		// 			}},
		// 		},
		// 		{
		// 			Category: to.Ptr("apiserver"),
		// 			Endpoints: []*armcontainerservice.EndpointDependency{
		// 				{
		// 					DomainName: to.Ptr("*.azmk8s.io"),
		// 					EndpointDetails: []*armcontainerservice.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 							Protocol: to.Ptr("Https"),
		// 					}},
		// 			}},
		// 		},
		// 		{
		// 			Category: to.Ptr("tunnel-classic"),
		// 			Endpoints: []*armcontainerservice.EndpointDependency{
		// 				{
		// 					DomainName: to.Ptr("*.azmk8s.io"),
		// 					EndpointDetails: []*armcontainerservice.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](9000),
		// 							Protocol: to.Ptr("TCP"),
		// 					}},
		// 			}},
		// 	}},
		// }
	}
}
