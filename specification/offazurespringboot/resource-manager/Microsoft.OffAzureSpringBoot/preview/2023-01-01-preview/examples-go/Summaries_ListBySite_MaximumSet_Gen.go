package armspringappdiscovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/springappdiscovery/armspringappdiscovery"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/41e4538ed7bb3ceac3c1322c9455a0812ed110ac/specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/Summaries_ListBySite_MaximumSet_Gen.json
func ExampleSummariesClient_NewListBySitePager_summariesListBySiteMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armspringappdiscovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSummariesClient().NewListBySitePager("rgspringbootdiscovery", "xxkzlvbihwxunadjcpjpjmghmhxrqyvghtpfps", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.SummaryList = armspringappdiscovery.SummaryList{
		// 	Value: []*armspringappdiscovery.Summary{
		// 		{
		// 			Name: to.Ptr("gntlsuizarosogjemzapadnhu"),
		// 			Type: to.Ptr("patxegcogfdxwobof"),
		// 			ID: to.Ptr("flpga"),
		// 			SystemData: &armspringappdiscovery.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-05T16:44:32.561Z"); return t}()),
		// 				CreatedBy: to.Ptr("ztjtyfhicmxcpqszeovgojwzceagbc"),
		// 				CreatedByType: to.Ptr(armspringappdiscovery.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-05T16:44:32.562Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("isjllzvqodp"),
		// 				LastModifiedByType: to.Ptr(armspringappdiscovery.CreatedByTypeUser),
		// 			},
		// 			Properties: &armspringappdiscovery.SummariesProperties{
		// 				DiscoveredApps: to.Ptr[int64](24),
		// 				DiscoveredServers: to.Ptr[int64](22),
		// 				ProvisioningState: to.Ptr(armspringappdiscovery.ProvisioningStateSucceeded),
		// 			},
		// 			Tags: map[string]*string{
		// 				"key9287": to.Ptr("utoxkd"),
		// 			},
		// 	}},
		// }
	}
}
