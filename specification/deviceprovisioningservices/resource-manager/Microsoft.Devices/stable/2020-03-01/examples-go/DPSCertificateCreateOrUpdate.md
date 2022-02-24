Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdeviceprovisioningservices%2Farmdeviceprovisioningservices%2Fv0.2.1/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armdeviceprovisioningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"
)

// x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2020-03-01/examples/DPSCertificateCreateOrUpdate.json
func ExampleDpsCertificateClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdeviceprovisioningservices.NewDpsCertificateClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<provisioning-service-name>",
		"<certificate-name>",
		armdeviceprovisioningservices.CertificateBodyDescription{
			Certificate: to.StringPtr("<certificate>"),
		},
		&armdeviceprovisioningservices.DpsCertificateClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DpsCertificateClientCreateOrUpdateResult)
}
```
