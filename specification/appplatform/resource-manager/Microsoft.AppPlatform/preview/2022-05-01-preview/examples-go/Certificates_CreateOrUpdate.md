Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fappplatform%2Farmappplatform%2Fv1.1.0-beta.1/sdk/resourcemanager/appplatform/armappplatform/README.md) on how to add the SDK to your project and authenticate.

```go
package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-05-01-preview/examples/Certificates_CreateOrUpdate.json
func ExampleCertificatesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappplatform.NewCertificatesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myservice",
		"mycertificate",
		armappplatform.CertificateResource{
			Properties: &armappplatform.KeyVaultCertificateProperties{
				Type:             to.Ptr("KeyVaultCertificate"),
				CertVersion:      to.Ptr("08a219d06d874795a96db47e06fbb01e"),
				KeyVaultCertName: to.Ptr("mycert"),
				VaultURI:         to.Ptr("https://myvault.vault.azure.net"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
