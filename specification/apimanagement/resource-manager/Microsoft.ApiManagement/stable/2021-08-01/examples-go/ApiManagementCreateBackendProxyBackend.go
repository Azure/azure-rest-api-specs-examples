package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateBackendProxyBackend.json
func ExampleBackendClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapimanagement.NewBackendClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"rg1",
		"apimService1",
		"proxybackend",
		armapimanagement.BackendContract{
			Properties: &armapimanagement.BackendContractProperties{
				Description: to.Ptr("description5308"),
				Credentials: &armapimanagement.BackendCredentialsContract{
					Authorization: &armapimanagement.BackendAuthorizationHeaderCredentials{
						Parameter: to.Ptr("opensesma"),
						Scheme:    to.Ptr("Basic"),
					},
					Header: map[string][]*string{
						"x-my-1": {
							to.Ptr("val1"),
							to.Ptr("val2")},
					},
					Query: map[string][]*string{
						"sv": {
							to.Ptr("xx"),
							to.Ptr("bb"),
							to.Ptr("cc")},
					},
				},
				Proxy: &armapimanagement.BackendProxyContract{
					Password: to.Ptr("<password>"),
					URL:      to.Ptr("http://192.168.1.1:8080"),
					Username: to.Ptr("Contoso\\admin"),
				},
				TLS: &armapimanagement.BackendTLSProperties{
					ValidateCertificateChain: to.Ptr(true),
					ValidateCertificateName:  to.Ptr(true),
				},
				URL:      to.Ptr("https://backendname2644/"),
				Protocol: to.Ptr(armapimanagement.BackendProtocolHTTP),
			},
		},
		&armapimanagement.BackendClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
