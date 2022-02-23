Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpolicyinsights%2Farmpolicyinsights%2Fv0.1.1/sdk/resourcemanager/policyinsights/armpolicyinsights/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_SummarizeSubscriptionLevelPolicySetDefinitionScope.json
func ExamplePolicyStatesClient_SummarizeForPolicySetDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicyinsights.NewPolicyStatesClient(cred, nil)
	res, err := client.SummarizeForPolicySetDefinition(ctx,
		armpolicyinsights.Enum6("latest"),
		"<subscription-id>",
		armpolicyinsights.Enum4("Microsoft.Authorization"),
		"<policy-set-definition-name>",
		&armpolicyinsights.QueryOptions{Top: to.Int32Ptr(1),
			Filter:    to.StringPtr("<filter>"),
			OrderBy:   nil,
			Select:    nil,
			From:      to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-05T18:00:00Z"); return t }()),
			To:        to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-06T18:00:00Z"); return t }()),
			Apply:     nil,
			SkipToken: nil,
			Expand:    nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PolicyStatesClientSummarizeForPolicySetDefinitionResult)
}
```
