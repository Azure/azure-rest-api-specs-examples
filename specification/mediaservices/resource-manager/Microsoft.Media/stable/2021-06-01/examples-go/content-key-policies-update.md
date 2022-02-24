Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmediaservices%2Farmmediaservices%2Fv0.3.1/sdk/resourcemanager/mediaservices/armmediaservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices"
)

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/content-key-policies-update.json
func ExampleContentKeyPoliciesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewContentKeyPoliciesClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<content-key-policy-name>",
		armmediaservices.ContentKeyPolicy{
			Properties: &armmediaservices.ContentKeyPolicyProperties{
				Description: to.StringPtr("<description>"),
				Options: []*armmediaservices.ContentKeyPolicyOption{
					{
						Name: to.StringPtr("<name>"),
						Configuration: &armmediaservices.ContentKeyPolicyClearKeyConfiguration{
							ODataType: to.StringPtr("<odata-type>"),
						},
						Restriction: &armmediaservices.ContentKeyPolicyOpenRestriction{
							ODataType: to.StringPtr("<odata-type>"),
						},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ContentKeyPoliciesClientUpdateResult)
}
```
