package armmigrate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/ImportCollectors_Create.json
func ExampleImportCollectorsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewImportCollectorsClient().Create(ctx, "markusavstestrg", "rajoshCCY9671project", "importCollector2952", &armmigrate.ImportCollectorsClientCreateOptions{CollectorBody: &armmigrate.ImportCollector{
		Name: to.Ptr("importCollector2951"),
		Type: to.Ptr("Microsoft.Migrate/assessmentprojects/importcollectors"),
		ETag: to.Ptr("\"000064a2-0000-3300-0000-605994800000\""),
		ID:   to.Ptr("/subscriptions/31be0ff4-c932-4cb3-8efc-efa411d79280/resourceGroups/markusavstestrg/providers/Microsoft.Migrate/assessmentprojects/rajoshCCY9671project/importcollectors/importCollector2951"),
		Properties: &armmigrate.ImportCollectorProperties{
			DiscoverySiteID: to.Ptr("/subscriptions/31be0ff4-c932-4cb3-8efc-efa411d79280/resourcegroups/MarkusAVStestRG/providers/microsoft.offazure/importsites/rajoshCCY54cbimportSite"),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ImportCollector = armmigrate.ImportCollector{
	// 	Name: to.Ptr("importCollector2952"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentprojects/importcollectors"),
	// 	ETag: to.Ptr("\"0000a7a2-0000-3300-0000-6059964d0000\""),
	// 	ID: to.Ptr("/subscriptions/31be0ff4-c932-4cb3-8efc-efa411d79280/resourceGroups/markusavstestrg/providers/Microsoft.Migrate/assessmentprojects/rajoshCCY9671project/importcollectors/importCollector2952"),
	// 	Properties: &armmigrate.ImportCollectorProperties{
	// 		CreatedTimestamp: to.Ptr("2021-03-23T07:18:37.2247735Z"),
	// 		DiscoverySiteID: to.Ptr("/subscriptions/31be0ff4-c932-4cb3-8efc-efa411d79280/resourcegroups/MarkusAVStestRG/providers/microsoft.offazure/importsites/rajoshCCY54cbimportSite"),
	// 		UpdatedTimestamp: to.Ptr("2021-03-23T07:18:37.2247735Z"),
	// 	},
	// }
}
