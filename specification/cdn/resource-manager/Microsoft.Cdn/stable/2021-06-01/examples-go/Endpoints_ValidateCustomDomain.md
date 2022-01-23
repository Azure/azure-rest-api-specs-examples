Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcdn%2Farmcdn%2Fv0.3.0/sdk/resourcemanager/cdn/armcdn/README.md) on how to add the SDK to your project and authenticate.

```go
package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Endpoints_ValidateCustomDomain.json
func ExampleEndpointsClient_ValidateCustomDomain() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcdn.NewEndpointsClient("<subscription-id>", cred, nil)
	res, err := client.ValidateCustomDomain(ctx,
		"<resource-group-name>",
		"<profile-name>",
		"<endpoint-name>",
		armcdn.ValidateCustomDomainInput{
			HostName: to.StringPtr("<host-name>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.EndpointsClientValidateCustomDomainResult)
}
```
