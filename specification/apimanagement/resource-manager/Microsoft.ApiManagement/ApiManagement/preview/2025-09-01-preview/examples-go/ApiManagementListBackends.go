package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementListBackends.json
func ExampleBackendClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBackendClient().NewListByServicePager("rg1", "apimService1", nil)
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
		// page = armapimanagement.BackendClientListByServiceResponse{
		// 	BackendCollection: armapimanagement.BackendCollection{
		// 		Count: to.Ptr[int64](4),
		// 		NextLink: to.Ptr(""),
		// 		Value: []*armapimanagement.BackendContract{
		// 			{
		// 				Name: to.Ptr("proxybackend"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/backends"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/backends/proxybackend"),
		// 				Properties: &armapimanagement.BackendContractProperties{
		// 					Description: to.Ptr("description5308"),
		// 					Credentials: &armapimanagement.BackendCredentialsContract{
		// 						Authorization: &armapimanagement.BackendAuthorizationHeaderCredentials{
		// 							Parameter: to.Ptr("opensesma"),
		// 							Scheme: to.Ptr("Basic"),
		// 						},
		// 						Header: map[string][]*string{
		// 							"x-my-1": []*string{
		// 								to.Ptr("val1"),
		// 								to.Ptr("val2"),
		// 							},
		// 						},
		// 						Query: map[string][]*string{
		// 							"sv": []*string{
		// 								to.Ptr("xx"),
		// 								to.Ptr("bb"),
		// 								to.Ptr("cc"),
		// 							},
		// 						},
		// 					},
		// 					Proxy: &armapimanagement.BackendProxyContract{
		// 						Password: to.Ptr("<password>"),
		// 						URL: to.Ptr("http://192.168.1.1:8080"),
		// 						Username: to.Ptr("Contoso\\admin"),
		// 					},
		// 					TLS: &armapimanagement.BackendTLSProperties{
		// 						ValidateCertificateChain: to.Ptr(false),
		// 						ValidateCertificateName: to.Ptr(false),
		// 					},
		// 					URL: to.Ptr("https://backendname2644/"),
		// 					Protocol: to.Ptr(armapimanagement.BackendProtocolHTTP),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("sfbackend"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/backends"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/backends/sfbackend"),
		// 				Properties: &armapimanagement.BackendContractProperties{
		// 					Description: to.Ptr("Service Fabric Test App 1"),
		// 					Properties: &armapimanagement.BackendProperties{
		// 						ServiceFabricCluster: &armapimanagement.BackendServiceFabricClusterProperties{
		// 							ClientCertificateID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/certificates/cert1"),
		// 							ManagementEndpoints: []*string{
		// 								to.Ptr("https://somecluster.com"),
		// 							},
		// 							MaxPartitionResolutionRetries: to.Ptr[int32](5),
		// 							ServerX509Names: []*armapimanagement.X509CertificateName{
		// 								{
		// 									Name: to.Ptr("ServerCommonName1"),
		// 									IssuerCertificateThumbprint: to.Ptr("IssuerCertificateThumbprint1"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					URL: to.Ptr("fabric:/mytestapp/mytestservice"),
		// 					Protocol: to.Ptr(armapimanagement.BackendProtocolHTTP),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("sfbackend"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/backends"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/backends/sustainable-backend-europe-north"),
		// 				Properties: &armapimanagement.BackendContractProperties{
		// 					Description: to.Ptr("description5308"),
		// 					AzureRegion: to.Ptr("northeurope"),
		// 					Credentials: &armapimanagement.BackendCredentialsContract{
		// 						Authorization: &armapimanagement.BackendAuthorizationHeaderCredentials{
		// 							Parameter: to.Ptr("opensesma"),
		// 							Scheme: to.Ptr("Basic"),
		// 						},
		// 						Header: map[string][]*string{
		// 							"x-my-1": []*string{
		// 								to.Ptr("val1"),
		// 								to.Ptr("val2"),
		// 							},
		// 						},
		// 						Query: map[string][]*string{
		// 							"sv": []*string{
		// 								to.Ptr("xx"),
		// 								to.Ptr("bb"),
		// 								to.Ptr("cc"),
		// 							},
		// 						},
		// 					},
		// 					Proxy: &armapimanagement.BackendProxyContract{
		// 						Password: to.Ptr("<password>"),
		// 						URL: to.Ptr("http://192.168.1.1:8080"),
		// 						Username: to.Ptr("Contoso\\admin"),
		// 					},
		// 					TLS: &armapimanagement.BackendTLSProperties{
		// 						ValidateCertificateChain: to.Ptr(false),
		// 						ValidateCertificateName: to.Ptr(false),
		// 					},
		// 					URL: to.Ptr("https://backendname2644/"),
		// 					Protocol: to.Ptr(armapimanagement.BackendProtocolHTTP),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("sfbackend"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/backends"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/backends/carbon-efficent-load-balancer"),
		// 				Properties: &armapimanagement.BackendContractProperties{
		// 					Type: to.Ptr(armapimanagement.BackendTypePool),
		// 					Pool: &armapimanagement.BackendBaseParametersPool{
		// 						Services: []*armapimanagement.BackendPoolItem{
		// 							{
		// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/backends/sustainable-backend-europe-north"),
		// 								PreferredCarbonEmission: to.Ptr(armapimanagement.CarbonEmissionCategoryMedium),
		// 								Priority: to.Ptr[int32](1),
		// 								Weight: to.Ptr[int32](1),
		// 							},
		// 							{
		// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/backends/proxybackend"),
		// 								Priority: to.Ptr[int32](1),
		// 								Weight: to.Ptr[int32](1),
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
