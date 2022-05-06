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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateGatewayHostnameConfiguration.json
func ExampleGatewayHostnameConfigurationClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armapimanagement.NewGatewayHostnameConfigurationClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<gateway-id>",
		"<hc-id>",
		armapimanagement.GatewayHostnameConfigurationContract{
			Properties: &armapimanagement.GatewayHostnameConfigurationContractProperties{
				CertificateID:              to.Ptr("<certificate-id>"),
				Hostname:                   to.Ptr("<hostname>"),
				HTTP2Enabled:               to.Ptr(true),
				NegotiateClientCertificate: to.Ptr(false),
				Tls10Enabled:               to.Ptr(false),
				Tls11Enabled:               to.Ptr(false),
			},
		},
		&armapimanagement.GatewayHostnameConfigurationClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
