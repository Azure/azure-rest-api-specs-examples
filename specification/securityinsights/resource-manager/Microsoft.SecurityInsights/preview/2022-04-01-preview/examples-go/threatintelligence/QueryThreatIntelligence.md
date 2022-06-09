```go
package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-04-01-preview/examples/threatintelligence/QueryThreatIntelligence.json
func ExampleThreatIntelligenceIndicatorClient_NewQueryIndicatorsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewThreatIntelligenceIndicatorClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.NewQueryIndicatorsPager("<resource-group-name>",
		"<workspace-name>",
		armsecurityinsights.ThreatIntelligenceFilteringCriteria{
			MaxConfidence: to.Ptr[int32](80),
			MaxValidUntil: to.Ptr("<max-valid-until>"),
			MinConfidence: to.Ptr[int32](25),
			MinValidUntil: to.Ptr("<min-valid-until>"),
			PageSize:      to.Ptr[int32](100),
			SortBy: []*armsecurityinsights.ThreatIntelligenceSortingCriteria{
				{
					ItemKey:   to.Ptr("<item-key>"),
					SortOrder: to.Ptr(armsecurityinsights.ThreatIntelligenceSortingCriteriaEnumDescending),
				}},
			Sources: []*string{
				to.Ptr("Azure Sentinel")},
		},
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurityinsights%2Farmsecurityinsights%2Fv0.3.0/sdk/resourcemanager/securityinsights/armsecurityinsights/README.md) on how to add the SDK to your project and authenticate.
