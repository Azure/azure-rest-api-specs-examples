Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsearch%2Farmsearch%2Fv0.3.0/sdk/resourcemanager/search/armsearch/README.md) on how to add the SDK to your project and authenticate.

```go
package armsearch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
)

// x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/SearchCheckNameAvailability.json
func ExampleServicesClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsearch.NewServicesClient("<subscription-id>", cred, nil)
	res, err := client.CheckNameAvailability(ctx,
		armsearch.CheckNameAvailabilityInput{
			Name: to.StringPtr("<name>"),
			Type: to.StringPtr("<type>"),
		},
		&armsearch.SearchManagementRequestOptions{ClientRequestID: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ServicesClientCheckNameAvailabilityResult)
}
```
