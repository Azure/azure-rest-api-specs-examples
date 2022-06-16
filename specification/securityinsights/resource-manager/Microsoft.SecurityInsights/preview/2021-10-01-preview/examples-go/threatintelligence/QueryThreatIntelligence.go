package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-10-01-preview/examples/threatintelligence/QueryThreatIntelligence.json
func ExampleThreatIntelligenceIndicatorClient_QueryIndicators() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsights.NewThreatIntelligenceIndicatorClient("<subscription-id>", cred, nil)
	pager := client.QueryIndicators("<resource-group-name>",
		"<workspace-name>",
		armsecurityinsights.ThreatIntelligenceFilteringCriteria{
			MaxConfidence: to.Int32Ptr(80),
			MaxValidUntil: to.StringPtr("<max-valid-until>"),
			MinConfidence: to.Int32Ptr(25),
			MinValidUntil: to.StringPtr("<min-valid-until>"),
			PageSize:      to.Int32Ptr(100),
			SortBy: []*armsecurityinsights.ThreatIntelligenceSortingCriteria{
				{
					ItemKey:   to.StringPtr("<item-key>"),
					SortOrder: armsecurityinsights.ThreatIntelligenceSortingCriteriaEnum("descending").ToPtr(),
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
