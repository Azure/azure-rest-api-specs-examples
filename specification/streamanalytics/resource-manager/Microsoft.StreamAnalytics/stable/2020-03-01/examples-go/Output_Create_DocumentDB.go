package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Create_DocumentDB.json
func ExampleOutputsClient_CreateOrReplace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstreamanalytics.NewOutputsClient("56b5e0a9-b645-407d-99b0-c64f86013e3d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrReplace(ctx,
		"sjrg7983",
		"sj2331",
		"output3022",
		armstreamanalytics.Output{
			Properties: &armstreamanalytics.OutputProperties{
				Datasource: &armstreamanalytics.DocumentDbOutputDataSource{
					Type: to.Ptr("Microsoft.Storage/DocumentDB"),
					Properties: &armstreamanalytics.DocumentDbOutputDataSourceProperties{
						AccountID:             to.Ptr("someAccountId"),
						AccountKey:            to.Ptr("accountKey=="),
						CollectionNamePattern: to.Ptr("collection"),
						Database:              to.Ptr("db01"),
						DocumentID:            to.Ptr("documentId"),
						PartitionKey:          to.Ptr("key"),
					},
				},
			},
		},
		&armstreamanalytics.OutputsClientCreateOrReplaceOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
