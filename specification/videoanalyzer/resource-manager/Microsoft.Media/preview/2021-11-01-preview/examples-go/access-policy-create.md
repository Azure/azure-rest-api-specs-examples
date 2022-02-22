Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fvideoanalyzer%2Farmvideoanalyzer%2Fv0.2.1/sdk/resourcemanager/videoanalyzer/armvideoanalyzer/README.md) on how to add the SDK to your project and authenticate.

```go
package armvideoanalyzer_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/videoanalyzer/armvideoanalyzer"
)

// x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/access-policy-create.json
func ExampleAccessPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvideoanalyzer.NewAccessPoliciesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<access-policy-name>",
		armvideoanalyzer.AccessPolicyEntity{
			Properties: &armvideoanalyzer.AccessPolicyProperties{
				Authentication: &armvideoanalyzer.JwtAuthentication{
					Type: to.StringPtr("<type>"),
					Audiences: []*string{
						to.StringPtr("audience1")},
					Claims: []*armvideoanalyzer.TokenClaim{
						{
							Name:  to.StringPtr("<name>"),
							Value: to.StringPtr("<value>"),
						},
						{
							Name:  to.StringPtr("<name>"),
							Value: to.StringPtr("<value>"),
						}},
					Issuers: []*string{
						to.StringPtr("issuer1"),
						to.StringPtr("issuer2")},
					Keys: []armvideoanalyzer.TokenKeyClassification{
						&armvideoanalyzer.RsaTokenKey{
							Type: to.StringPtr("<type>"),
							Kid:  to.StringPtr("<kid>"),
							Alg:  armvideoanalyzer.AccessPolicyRsaAlgo("RS256").ToPtr(),
							E:    to.StringPtr("<e>"),
							N:    to.StringPtr("<n>"),
						},
						&armvideoanalyzer.EccTokenKey{
							Type: to.StringPtr("<type>"),
							Kid:  to.StringPtr("<kid>"),
							Alg:  armvideoanalyzer.AccessPolicyEccAlgo("ES256").ToPtr(),
							X:    to.StringPtr("<x>"),
							Y:    to.StringPtr("<y>"),
						}},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AccessPoliciesClientCreateOrUpdateResult)
}
```
