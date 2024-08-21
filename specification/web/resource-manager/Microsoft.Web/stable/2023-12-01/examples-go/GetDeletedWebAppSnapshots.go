package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/21c2852d62ccc3abe9cc3800c989c6826f8363dc/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/GetDeletedWebAppSnapshots.json
func ExampleGlobalClient_GetDeletedWebAppSnapshots() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.SnapshotArray = []*armappservice.Snapshot{
	// 	{
	// 		Name: to.Ptr("wussite6"),
	// 		Type: to.Ptr("Microsoft.Web/locations/deletedSites"),
	// 		ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg1/providers/Microsoft.Web/locations/West US 2/deletedwebapps/wussite6/snapshots/9"),
	// 		Properties: &armappservice.SnapshotProperties{
	// 			Time: to.Ptr("2019-05-09T22:29:05.1337007"),
	// 		},
	// }}
}
