Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpolicyinsights%2Farmpolicyinsights%2Fv0.4.0/sdk/resourcemanager/policyinsights/armpolicyinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armpolicyinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_SummarizeSubscriptionLevelPolicySetDefinitionScope.json
func ExamplePolicyStatesClient_SummarizeForPolicySetDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armpolicyinsights.NewPolicyStatesClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.SummarizeForPolicySetDefinition(ctx,
		armpolicyinsights.PolicyStatesSummaryResourceTypeLatest,
		"<subscription-id>",
		"<policy-set-definition-name>",
		&armpolicyinsights.QueryOptions{Top: to.Ptr[int32](1),
			Filter:    to.Ptr("<filter>"),
			OrderBy:   nil,
			Select:    nil,
			From:      to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-05T18:00:00Z"); return t }()),
			To:        to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-06T18:00:00Z"); return t }()),
			Apply:     nil,
			SkipToken: nil,
			Expand:    nil,
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
