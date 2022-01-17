Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpolicyinsights%2Farmpolicyinsights%2Fv0.1.0/sdk/resourcemanager/policyinsights/armpolicyinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_SummarizeSubscriptionLevelPolicyAssignmentScope.json
func ExamplePolicyStatesClient_SummarizeForSubscriptionLevelPolicyAssignment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicyinsights.NewPolicyStatesClient(cred, nil)
	res, err := client.SummarizeForSubscriptionLevelPolicyAssignment(ctx,
		armpolicyinsights.Enum6("latest"),
		"<subscription-id>",
		armpolicyinsights.Enum4("Microsoft.Authorization"),
		"<policy-assignment-name>",
		&armpolicyinsights.QueryOptions{Top: nil,
			Filter:    nil,
			OrderBy:   nil,
			Select:    nil,
			From:      nil,
			To:        nil,
			Apply:     nil,
			SkipToken: nil,
			Expand:    nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PolicyStatesClientSummarizeForSubscriptionLevelPolicyAssignmentResult)
}
```
