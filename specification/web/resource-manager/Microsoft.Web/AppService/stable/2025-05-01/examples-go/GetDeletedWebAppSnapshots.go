package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v6"
)

// Generated from example definition: 2025-05-01/GetDeletedWebAppSnapshots.json
func ExampleGlobalClient_GetDeletedWebAppSnapshots() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGlobalClient().GetDeletedWebAppSnapshots(ctx, "9", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armappservice.GlobalClientGetDeletedWebAppSnapshotsResponse{
	// 	undefined: &[]*armappservice.Snapshot{
	// 		{
	// 			Name: to.Ptr("wussite6"),
	// 			Type: to.Ptr("Microsoft.Web/locations/deletedSites"),
	// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg1/providers/Microsoft.Web/locations/West US 2/deletedwebapps/wussite6/snapshots/9"),
	// 			Properties: &armappservice.SnapshotProperties{
	// 				Time: to.Ptr("2019-05-09T22:29:05.1337007"),
	// 			},
	// 		},
	// 	},
	// }
}
