Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpolicyinsights%2Farmpolicyinsights%2Fv0.4.0/sdk/resourcemanager/policyinsights/armpolicyinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-03-01/examples/PolicyRestrictions_CheckAtSubscriptionScope.json
func ExamplePolicyRestrictionsClient_CheckAtSubscriptionScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armpolicyinsights.NewPolicyRestrictionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CheckAtSubscriptionScope(ctx,
		armpolicyinsights.CheckRestrictionsRequest{
			PendingFields: []*armpolicyinsights.PendingField{
				{
					Field: to.Ptr("<field>"),
					Values: []*string{
						to.Ptr("myVMName")},
				},
				{
					Field: to.Ptr("<field>"),
					Values: []*string{
						to.Ptr("eastus"),
						to.Ptr("westus"),
						to.Ptr("westus2"),
						to.Ptr("westeurope")},
				},
				{
					Field: to.Ptr("<field>"),
				}},
			ResourceDetails: &armpolicyinsights.CheckRestrictionsResourceDetails{
				APIVersion: to.Ptr("<apiversion>"),
				ResourceContent: map[string]interface{}{
					"type": "Microsoft.Compute/virtualMachines",
					"properties": map[string]interface{}{
						"priority": "Spot",
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
