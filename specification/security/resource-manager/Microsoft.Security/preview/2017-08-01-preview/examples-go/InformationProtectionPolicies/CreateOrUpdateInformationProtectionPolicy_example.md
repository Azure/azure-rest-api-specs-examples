Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.6.0/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/InformationProtectionPolicies/CreateOrUpdateInformationProtectionPolicy_example.json
func ExampleInformationProtectionPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsecurity.NewInformationProtectionPoliciesClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<scope>",
		armsecurity.InformationProtectionPolicyNameCustom,
		armsecurity.InformationProtectionPolicy{
			Properties: &armsecurity.InformationProtectionPolicyProperties{
				InformationTypes: map[string]*armsecurity.InformationType{
					"3bf35491-99b8-41f2-86d5-c1200a7df658": {
						Custom:      to.Ptr(true),
						DisplayName: to.Ptr("<display-name>"),
						Enabled:     to.Ptr(true),
						Keywords: []*armsecurity.InformationProtectionKeyword{
							{
								CanBeNumeric: to.Ptr(true),
								Custom:       to.Ptr(true),
								Pattern:      to.Ptr("<pattern>"),
							}},
						Order:              to.Ptr[int32](1400),
						RecommendedLabelID: to.Ptr("<recommended-label-id>"),
					},
					"7fb9419d-2473-4ad8-8e11-b25cc8cf6a07": {
						Custom:      to.Ptr(false),
						DisplayName: to.Ptr("<display-name>"),
						Enabled:     to.Ptr(true),
						Keywords: []*armsecurity.InformationProtectionKeyword{
							{
								CanBeNumeric: to.Ptr(false),
								Custom:       to.Ptr(true),
								Pattern:      to.Ptr("<pattern>"),
							}},
						Order:              to.Ptr[int32](100),
						RecommendedLabelID: to.Ptr("<recommended-label-id>"),
					},
				},
				Labels: map[string]*armsecurity.SensitivityLabel{
					"1345da73-bc5a-4a8f-b7dd-3820eb713da8": {
						DisplayName: to.Ptr("<display-name>"),
						Enabled:     to.Ptr(true),
						Order:       to.Ptr[int32](100),
					},
					"575739d2-3d53-4df0-9042-4c7772d5c7b1": {
						DisplayName: to.Ptr("<display-name>"),
						Enabled:     to.Ptr(true),
						Order:       to.Ptr[int32](300),
					},
					"7aa516c7-5a53-4857-bc6e-6808c6acd542": {
						DisplayName: to.Ptr("<display-name>"),
						Enabled:     to.Ptr(true),
						Order:       to.Ptr[int32](200),
					},
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
