Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Ffrontdoor%2Farmfrontdoor%2Fv0.1.0/sdk/resourcemanager/frontdoor/armfrontdoor/README.md) on how to add the SDK to your project and authenticate.

```go
package armfrontdoor_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorEnableHttps.json
func ExampleFrontendEndpointsClient_BeginEnableHTTPS() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armfrontdoor.NewFrontendEndpointsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginEnableHTTPS(ctx,
		"<resource-group-name>",
		"<front-door-name>",
		"<frontend-endpoint-name>",
		armfrontdoor.CustomHTTPSConfiguration{
			CertificateSource: armfrontdoor.FrontDoorCertificateSourceAzureKeyVault.ToPtr(),
			KeyVaultCertificateSourceParameters: &armfrontdoor.KeyVaultCertificateSourceParameters{
				SecretName:    to.StringPtr("<secret-name>"),
				SecretVersion: to.StringPtr("<secret-version>"),
				Vault: &armfrontdoor.KeyVaultCertificateSourceParametersVault{
					ID: to.StringPtr("<id>"),
				},
			},
			MinimumTLSVersion: armfrontdoor.MinimumTLSVersionOne0.ToPtr(),
			ProtocolType:      armfrontdoor.FrontDoorTLSProtocolTypeServerNameIndication.ToPtr(),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
```
