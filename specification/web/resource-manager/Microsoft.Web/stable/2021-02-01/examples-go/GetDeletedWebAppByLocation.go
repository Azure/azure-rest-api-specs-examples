package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/GetDeletedWebAppByLocation.json
func ExampleDeletedWebAppsClient_GetDeletedWebAppByLocation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewDeletedWebAppsClient("<subscription-id>", cred, nil)
	res, err := client.GetDeletedWebAppByLocation(ctx,
		"<location>",
		"<deleted-site-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DeletedSite.ID: %s\n", *res.ID)
}
