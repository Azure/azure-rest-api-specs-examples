package armintegrationspaces_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/integrationspaces/armintegrationspaces"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/azureintegrationspaces/resource-manager/Microsoft.IntegrationSpaces/preview/2023-11-14-preview/examples/Applications_ListBySpace.json
func ExampleApplicationsClient_NewListBySpacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armintegrationspaces.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewApplicationsClient().NewListBySpacePager("testrg", "Space1", &armintegrationspaces.ApplicationsClientListBySpaceOptions{Top: nil,
		Skip:        nil,
		Maxpagesize: nil,
		Filter:      nil,
		Select:      []string{},
		Expand:      []string{},
		Orderby:     []string{},
	})
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
		// page.ApplicationListResult = armintegrationspaces.ApplicationListResult{
		// 	Value: []*armintegrationspaces.Application{
		// 		{
		// 			Name: to.Ptr("Application1"),
		// 			Type: to.Ptr("Microsoft.IntegrationSpaces/spaces/applications"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.IntegrationSpaces/spaces/Space1/applications/Application1"),
		// 			Location: to.Ptr("CentralUS"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("Value1"),
		// 			},
		// 			Properties: &armintegrationspaces.ApplicationProperties{
		// 				Description: to.Ptr("This is the user provided description of the application."),
		// 				ProvisioningState: to.Ptr(armintegrationspaces.ProvisioningStateSucceeded),
		// 				TrackingDataStores: map[string]*armintegrationspaces.TrackingDataStore{
		// 					"dataStoreName1": &armintegrationspaces.TrackingDataStore{
		// 						DataStoreIngestionURI: to.Ptr("https://ingest-someClusterName.someRegionName.kusto.windows.net"),
		// 						DataStoreResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Kusto/Clusters/cluster1"),
		// 						DataStoreURI: to.Ptr("https://someClusterName.someRegionName.kusto.windows.net"),
		// 						DatabaseName: to.Ptr("testDatabase1"),
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
