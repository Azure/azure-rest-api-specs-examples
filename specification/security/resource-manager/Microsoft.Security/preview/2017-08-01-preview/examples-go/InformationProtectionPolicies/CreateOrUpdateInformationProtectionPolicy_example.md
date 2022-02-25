Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.4.0/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/InformationProtectionPolicies/CreateOrUpdateInformationProtectionPolicy_example.json
func ExampleInformationProtectionPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewInformationProtectionPoliciesClient(cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<scope>",
		armsecurity.InformationProtectionPolicyName("custom"),
		armsecurity.InformationProtectionPolicy{
			Properties: &armsecurity.InformationProtectionPolicyProperties{
				InformationTypes: map[string]*armsecurity.InformationType{
					"3bf35491-99b8-41f2-86d5-c1200a7df658": {
						Custom:      to.BoolPtr(true),
						DisplayName: to.StringPtr("<display-name>"),
						Enabled:     to.BoolPtr(true),
						Keywords: []*armsecurity.InformationProtectionKeyword{
							{
								CanBeNumeric: to.BoolPtr(true),
								Custom:       to.BoolPtr(true),
								Pattern:      to.StringPtr("<pattern>"),
							}},
						Order:              to.Int32Ptr(1400),
						RecommendedLabelID: to.StringPtr("<recommended-label-id>"),
					},
					"7fb9419d-2473-4ad8-8e11-b25cc8cf6a07": {
						Custom:      to.BoolPtr(false),
						DisplayName: to.StringPtr("<display-name>"),
						Enabled:     to.BoolPtr(true),
						Keywords: []*armsecurity.InformationProtectionKeyword{
							{
								CanBeNumeric: to.BoolPtr(false),
								Custom:       to.BoolPtr(true),
								Pattern:      to.StringPtr("<pattern>"),
							}},
						Order:              to.Int32Ptr(100),
						RecommendedLabelID: to.StringPtr("<recommended-label-id>"),
					},
				},
				Labels: map[string]*armsecurity.SensitivityLabel{
					"1345da73-bc5a-4a8f-b7dd-3820eb713da8": {
						DisplayName: to.StringPtr("<display-name>"),
						Enabled:     to.BoolPtr(true),
						Order:       to.Int32Ptr(100),
					},
					"575739d2-3d53-4df0-9042-4c7772d5c7b1": {
						DisplayName: to.StringPtr("<display-name>"),
						Enabled:     to.BoolPtr(true),
						Order:       to.Int32Ptr(300),
					},
					"7aa516c7-5a53-4857-bc6e-6808c6acd542": {
						DisplayName: to.StringPtr("<display-name>"),
						Enabled:     to.BoolPtr(true),
						Order:       to.Int32Ptr(200),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.InformationProtectionPoliciesClientCreateOrUpdateResult)
}
```
