package armdatabox_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/databox/resource-manager/Microsoft.DataBox/stable/2022-02-01/examples/BookShipmentPickupPost.json
func ExampleJobsClient_BookShipmentPickUp() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatabox.NewJobsClient("fa68082f-8ff7-4a25-95c7-ce9da541242f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.BookShipmentPickUp(ctx,
		"bvttoolrg6",
		"TJ-636646322037905056",
		armdatabox.ShipmentPickUpRequest{
			EndTime:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-22T18:30:00Z"); return t }()),
			ShipmentLocation: to.Ptr("Front desk"),
			StartTime:        to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-20T18:30:00Z"); return t }()),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
