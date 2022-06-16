package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/threatintelligence/QueryThreatIntelligence.json
func ExampleThreatIntelligenceIndicatorClient_NewQueryIndicatorsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewThreatIntelligenceIndicatorClient("bd794837-4d29-4647-9105-6339bfdb4e6a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewQueryIndicatorsPager("myRg",
		"myWorkspace",
		armsecurityinsights.ThreatIntelligenceFilteringCriteria{
			MaxConfidence: to.Ptr[int32](80),
			MaxValidUntil: to.Ptr("2020-04-25T17:44:00.114052Z"),
			MinConfidence: to.Ptr[int32](25),
			MinValidUntil: to.Ptr("2020-04-05T17:44:00.114052Z"),
			PageSize:      to.Ptr[int32](100),
			SortBy: []*armsecurityinsights.ThreatIntelligenceSortingCriteria{
				{
					ItemKey:   to.Ptr("lastUpdatedTimeUtc"),
					SortOrder: to.Ptr(armsecurityinsights.ThreatIntelligenceSortingOrderDescending),
				}},
			Sources: []*string{
				to.Ptr("Azure Sentinel")},
		},
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
