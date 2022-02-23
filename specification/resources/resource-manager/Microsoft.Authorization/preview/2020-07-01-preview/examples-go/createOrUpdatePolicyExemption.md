Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresources%2Farmpolicy%2Fv0.3.1/sdk/resourcemanager/resources/armpolicy/README.md) on how to add the SDK to your project and authenticate.

```go
package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2020-07-01-preview/examples/createOrUpdatePolicyExemption.json
func ExampleExemptionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicy.NewExemptionsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<scope>",
		"<policy-exemption-name>",
		armpolicy.Exemption{
			Properties: &armpolicy.ExemptionProperties{
				Description:       to.StringPtr("<description>"),
				DisplayName:       to.StringPtr("<display-name>"),
				ExemptionCategory: armpolicy.ExemptionCategory("Waiver").ToPtr(),
				Metadata: map[string]interface{}{
					"reason": "Temporary exemption for a expensive VM demo",
				},
				PolicyAssignmentID: to.StringPtr("<policy-assignment-id>"),
				PolicyDefinitionReferenceIDs: []*string{
					to.StringPtr("Limit_Skus")},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ExemptionsClientCreateOrUpdateResult)
}
```
