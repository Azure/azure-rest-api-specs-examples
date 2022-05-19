Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurityinsights%2Farmsecurityinsights%2Fv1.0.0/sdk/resourcemanager/securityinsights/armsecurityinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/threatintelligence/ReplaceTagsThreatIntelligence.json
func ExampleThreatIntelligenceIndicatorClient_ReplaceTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewThreatIntelligenceIndicatorClient("bd794837-4d29-4647-9105-6339bfdb4e6a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ReplaceTags(ctx,
		"myRg",
		"myWorkspace",
		"d9cd6f0b-96b9-3984-17cd-a779d1e15a93",
		armsecurityinsights.ThreatIntelligenceIndicatorModel{
			Etag: to.Ptr("\"0000262c-0000-0800-0000-5e9767060000\""),
			Kind: to.Ptr(armsecurityinsights.ThreatIntelligenceResourceInnerKindIndicator),
			Properties: &armsecurityinsights.ThreatIntelligenceIndicatorProperties{
				ThreatIntelligenceTags: []*string{
					to.Ptr("patching tags")},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
