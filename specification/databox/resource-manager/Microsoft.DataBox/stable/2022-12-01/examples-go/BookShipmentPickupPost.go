package armdatabox_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/BookShipmentPickupPost.json
func ExampleJobsClient_BookShipmentPickUp() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabox.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobsClient().BookShipmentPickUp(ctx, "YourResourceGroupName", "TestJobName1", armdatabox.ShipmentPickUpRequest{
		EndTime:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-22T18:30:00.000Z"); return t }()),
		ShipmentLocation: to.Ptr("Front desk"),
		StartTime:        to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-20T18:30:00.000Z"); return t }()),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ShipmentPickUpResponse = armdatabox.ShipmentPickUpResponse{
	// 	ConfirmationNumber: to.Ptr("XXXXXXXXXXX"),
	// 	ReadyByTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-20T18:30:00.000Z"); return t}()),
	// }
}
