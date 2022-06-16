package armmigrate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/ImportCollectors_Create.json
func ExampleImportCollectorsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmigrate.NewImportCollectorsClient("31be0ff4-c932-4cb3-8efc-efa411d79280", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"markusavstestrg",
		"rajoshCCY9671project",
		"importCollector2952",
		&armmigrate.ImportCollectorsClientCreateOptions{CollectorBody: &armmigrate.ImportCollector{
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
	// TODO: use response item
	_ = res
}
