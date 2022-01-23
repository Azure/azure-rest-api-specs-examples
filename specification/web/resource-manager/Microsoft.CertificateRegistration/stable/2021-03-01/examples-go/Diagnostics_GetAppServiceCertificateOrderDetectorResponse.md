Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fappservice%2Farmappservice%2Fv0.2.0/sdk/resourcemanager/appservice/armappservice/README.md) on how to add the SDK to your project and authenticate.

```go
package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2021-03-01/examples/Diagnostics_GetAppServiceCertificateOrderDetectorResponse.json
func ExampleCertificateOrdersDiagnosticsClient_GetAppServiceCertificateOrderDetectorResponse() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewCertificateOrdersDiagnosticsClient("<subscription-id>", cred, nil)
	res, err := client.GetAppServiceCertificateOrderDetectorResponse(ctx,
		"<resource-group-name>",
		"<certificate-order-name>",
		"<detector-name>",
		&armappservice.CertificateOrdersDiagnosticsClientGetAppServiceCertificateOrderDetectorResponseOptions{StartTime: nil,
			EndTime:   nil,
			TimeGrain: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.CertificateOrdersDiagnosticsClientGetAppServiceCertificateOrderDetectorResponseResult)
}
```
