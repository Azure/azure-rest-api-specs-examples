```go
package armsecurityinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsight/armsecurityinsight"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/threatintelligence/QueryThreatIntelligence.json
func ExampleThreatIntelligenceIndicatorClient_QueryIndicators() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsight.NewThreatIntelligenceIndicatorClient("<subscription-id>", cred, nil)
	pager := client.QueryIndicators("<resource-group-name>",
		"<workspace-name>",
		armsecurityinsight.ThreatIntelligenceFilteringCriteria{
			MaxConfidence: to.Int32Ptr(80),
			MaxValidUntil: to.StringPtr("<max-valid-until>"),
			MinConfidence: to.Int32Ptr(25),
			MinValidUntil: to.StringPtr("<min-valid-until>"),
			PageSize:      to.Int32Ptr(100),
			SortBy: []*armsecurityinsight.ThreatIntelligenceSortingCriteria{
				{
					ItemKey:   to.StringPtr("<item-key>"),
					SortOrder: armsecurityinsight.ThreatIntelligenceSortingCriteriaEnum("descending").ToPtr(),
				}},
			Sources: []*string{
				to.StringPtr("Azure Sentinel")},
		},
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurityinsight%2Farmsecurityinsight%2Fv0.2.1/sdk/resourcemanager/securityinsight/armsecurityinsight/README.md) on how to add the SDK to your project and authenticate.
