Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmediaservices%2Farmmediaservices%2Fv0.6.0/sdk/resourcemanager/mediaservices/armmediaservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/content-key-policies-create-nodrm-token.json
func ExampleContentKeyPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmediaservices.NewContentKeyPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<content-key-policy-name>",
		armmediaservices.ContentKeyPolicy{
			Properties: &armmediaservices.ContentKeyPolicyProperties{
				Description: to.Ptr("<description>"),
				Options: []*armmediaservices.ContentKeyPolicyOption{
					{
						Name: to.Ptr("<name>"),
						Configuration: &armmediaservices.ContentKeyPolicyClearKeyConfiguration{
							ODataType: to.Ptr("<odata-type>"),
						},
						Restriction: &armmediaservices.ContentKeyPolicyTokenRestriction{
							ODataType: to.Ptr("<odata-type>"),
							Audience:  to.Ptr("<audience>"),
							Issuer:    to.Ptr("<issuer>"),
							PrimaryVerificationKey: &armmediaservices.ContentKeyPolicySymmetricTokenKey{
								ODataType: to.Ptr("<odata-type>"),
								KeyValue:  []byte("AAAAAAAAAAAAAAAAAAAAAA=="),
							},
							RestrictionTokenType: to.Ptr(armmediaservices.ContentKeyPolicyRestrictionTokenTypeSwt),
						},
					}},
			},
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
