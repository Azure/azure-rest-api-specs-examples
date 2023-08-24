package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateBackend.json
func ExampleBackendClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBackendClient().Update(ctx, "rg1", "apimService1", "proxybackend", "*", armapimanagement.BackendUpdateParameters{
		Properties: &armapimanagement.BackendUpdateParameterProperties{
			Description: to.Ptr("description5308"),
			TLS: &armapimanagement.BackendTLSProperties{
				ValidateCertificateChain: to.Ptr(false),
				ValidateCertificateName:  to.Ptr(true),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackendContract = armapimanagement.BackendContract{
	// 	Name: to.Ptr("proxybackend"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/backends"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/backends/proxybackend"),
	// 	Properties: &armapimanagement.BackendContractProperties{
	// 		Description: to.Ptr("description5308"),
	// 		Credentials: &armapimanagement.BackendCredentialsContract{
	// 			Authorization: &armapimanagement.BackendAuthorizationHeaderCredentials{
	// 				Parameter: to.Ptr("opensesma"),
	// 				Scheme: to.Ptr("Basic"),
	// 			},
	// 			Header: map[string][]*string{
	// 				"x-my-1": []*string{
	// 					to.Ptr("val1"),
	// 					to.Ptr("val2")},
	// 				},
	// 				Query: map[string][]*string{
	// 					"sv": []*string{
	// 						to.Ptr("xx"),
	// 						to.Ptr("bb"),
	// 						to.Ptr("cc")},
	// 					},
	// 				},
	// 				Proxy: &armapimanagement.BackendProxyContract{
	// 					Password: to.Ptr("<password>"),
	// 					URL: to.Ptr("http://192.168.1.1:8080"),
	// 					Username: to.Ptr("Contoso\\admin"),
	// 				},
	// 				TLS: &armapimanagement.BackendTLSProperties{
	// 					ValidateCertificateChain: to.Ptr(false),
	// 					ValidateCertificateName: to.Ptr(true),
	// 				},
	// 				URL: to.Ptr("https://backendname2644/"),
	// 				Protocol: to.Ptr(armapimanagement.BackendProtocolHTTP),
	// 			},
	// 		}
}
