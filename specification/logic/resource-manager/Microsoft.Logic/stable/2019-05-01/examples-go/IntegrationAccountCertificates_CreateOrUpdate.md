Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Flogic%2Farmlogic%2Fv0.3.1/sdk/resourcemanager/logic/armlogic/README.md) on how to add the SDK to your project and authenticate.

```go
package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountCertificates_CreateOrUpdate.json
func ExampleIntegrationAccountCertificatesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogic.NewIntegrationAccountCertificatesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<integration-account-name>",
		"<certificate-name>",
		armlogic.IntegrationAccountCertificate{
			Location: to.StringPtr("<location>"),
			Properties: &armlogic.IntegrationAccountCertificateProperties{
				Key: &armlogic.KeyVaultKeyReference{
					KeyName: to.StringPtr("<key-name>"),
					KeyVault: &armlogic.KeyVaultKeyReferenceKeyVault{
						ID: to.StringPtr("<id>"),
					},
					KeyVersion: to.StringPtr("<key-version>"),
				},
				PublicCertificate: to.StringPtr("<public-certificate>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IntegrationAccountCertificatesClientCreateOrUpdateResult)
}
```
