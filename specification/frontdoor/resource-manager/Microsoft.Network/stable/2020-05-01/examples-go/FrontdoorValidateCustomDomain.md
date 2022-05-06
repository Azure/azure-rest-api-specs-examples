Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Ffrontdoor%2Farmfrontdoor%2Fv0.4.0/sdk/resourcemanager/frontdoor/armfrontdoor/README.md) on how to add the SDK to your project and authenticate.

```go
package armfrontdoor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorValidateCustomDomain.json
func ExampleFrontDoorsClient_ValidateCustomDomain() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armfrontdoor.NewFrontDoorsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.ValidateCustomDomain(ctx,
		"<resource-group-name>",
		"<front-door-name>",
		armfrontdoor.ValidateCustomDomainInput{
			HostName: to.Ptr("<host-name>"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
