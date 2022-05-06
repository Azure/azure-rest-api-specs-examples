Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fapimanagement%2Farmapimanagement%2Fv0.5.0/sdk/resourcemanager/apimanagement/armapimanagement/README.md) on how to add the SDK to your project and authenticate.

```go
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
		return
	}
	ctx := context.Background()
	client, err := armapimanagement.NewBackendClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<backend-id>",
		armapimanagement.BackendContract{
			Properties: &armapimanagement.BackendContractProperties{
				Description: to.Ptr("<description>"),
				Credentials: &armapimanagement.BackendCredentialsContract{
					Authorization: &armapimanagement.BackendAuthorizationHeaderCredentials{
						Parameter: to.Ptr("<parameter>"),
						Scheme:    to.Ptr("<scheme>"),
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
					URL:      to.Ptr("<url>"),
					Username: to.Ptr("<username>"),
				},
				TLS: &armapimanagement.BackendTLSProperties{
					ValidateCertificateChain: to.Ptr(true),
					ValidateCertificateName:  to.Ptr(true),
				},
				URL:      to.Ptr("<url>"),
				Protocol: to.Ptr(armapimanagement.BackendProtocolHTTP),
			},
		},
		&armapimanagement.BackendClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
