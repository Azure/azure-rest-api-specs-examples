package armmigrate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/ImportCollectors_ListByProject.json
func ExampleImportCollectorsClient_NewListByProjectPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewImportCollectorsClient().NewListByProjectPager("markusavstestrg", "rajoshCCY9671project", nil)
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
		// page.ImportCollectorList = armmigrate.ImportCollectorList{
		// 	Value: []*armmigrate.ImportCollector{
		// 		{
		// 			Name: to.Ptr("importCollector2951"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects/importcollectors"),
		// 			ETag: to.Ptr("\"000098a2-0000-3300-0000-605995620000\""),
		// 			ID: to.Ptr("/subscriptions/31be0ff4-c932-4cb3-8efc-efa411d79280/resourceGroups/markusavstestrg/providers/Microsoft.Migrate/assessmentprojects/rajoshCCY9671project/importcollectors/importCollector2951"),
		// 			Properties: &armmigrate.ImportCollectorProperties{
		// 				CreatedTimestamp: to.Ptr("2021-02-11T04:46:54.9582099Z"),
		// 				DiscoverySiteID: to.Ptr("/subscriptions/31be0ff4-c932-4cb3-8efc-efa411d79280/resourcegroups/MarkusAVStestRG/providers/microsoft.offazure/importsites/rajoshCCY54cbimportSite"),
		// 				UpdatedTimestamp: to.Ptr("2021-03-23T07:14:42.9238657Z"),
		// 			},
		// 	}},
		// }
	}
}
