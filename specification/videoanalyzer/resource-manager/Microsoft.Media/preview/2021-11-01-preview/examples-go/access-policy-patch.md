Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fvideoanalyzer%2Farmvideoanalyzer%2Fv0.4.0/sdk/resourcemanager/videoanalyzer/armvideoanalyzer/README.md) on how to add the SDK to your project and authenticate.

```go
package armvideoanalyzer_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/videoanalyzer/armvideoanalyzer"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/access-policy-patch.json
func ExampleAccessPoliciesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armvideoanalyzer.NewAccessPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<access-policy-name>",
		armvideoanalyzer.AccessPolicyEntity{
			Properties: &armvideoanalyzer.AccessPolicyProperties{
				Authentication: &armvideoanalyzer.JwtAuthentication{
					Type: to.Ptr("<type>"),
					Audiences: []*string{
						to.Ptr("audience1")},
					Claims: []*armvideoanalyzer.TokenClaim{
						{
							Name:  to.Ptr("<name>"),
							Value: to.Ptr("<value>"),
						},
						{
							Name:  to.Ptr("<name>"),
							Value: to.Ptr("<value>"),
						}},
					Issuers: []*string{
						to.Ptr("issuer1"),
						to.Ptr("issuer2")},
					Keys: []armvideoanalyzer.TokenKeyClassification{
						&armvideoanalyzer.RsaTokenKey{
							Type: to.Ptr("<type>"),
							Kid:  to.Ptr("<kid>"),
							Alg:  to.Ptr(armvideoanalyzer.AccessPolicyRsaAlgoRS256),
							E:    to.Ptr("<e>"),
							N:    to.Ptr("<n>"),
						},
						&armvideoanalyzer.EccTokenKey{
							Type: to.Ptr("<type>"),
							Kid:  to.Ptr("<kid>"),
							Alg:  to.Ptr(armvideoanalyzer.AccessPolicyEccAlgoES256),
							X:    to.Ptr("<x>"),
							Y:    to.Ptr("<y>"),
						}},
				},
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
