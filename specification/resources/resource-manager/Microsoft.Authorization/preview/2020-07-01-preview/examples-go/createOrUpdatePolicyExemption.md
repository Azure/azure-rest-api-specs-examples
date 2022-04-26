Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresources%2Farmpolicy%2Fv0.5.0/sdk/resourcemanager/resources/armpolicy/README.md) on how to add the SDK to your project and authenticate.

```go
package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Authorization/preview/2020-07-01-preview/examples/createOrUpdatePolicyExemption.json
func ExampleExemptionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armpolicy.NewExemptionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<scope>",
		"<policy-exemption-name>",
		armpolicy.Exemption{
			Properties: &armpolicy.ExemptionProperties{
				Description:       to.Ptr("<description>"),
				DisplayName:       to.Ptr("<display-name>"),
				ExemptionCategory: to.Ptr(armpolicy.ExemptionCategoryWaiver),
				Metadata: map[string]interface{}{
					"reason": "Temporary exemption for a expensive VM demo",
				},
				PolicyAssignmentID: to.Ptr("<policy-assignment-id>"),
				PolicyDefinitionReferenceIDs: []*string{
					to.Ptr("Limit_Skus")},
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
