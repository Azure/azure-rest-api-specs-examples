package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Accounts/stable/2021-11-01/examples/media-service-operation-status-by-id-terminal-state.json
func ExampleOperationStatusesClient_Get_getStatusOfAsynchronousOperationWhenItIsCompleted() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationStatusesClient().Get(ctx, "westus", "D612C429-2526-49D5-961B-885AE11406FD", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MediaServiceOperationStatus = armmediaservices.MediaServiceOperationStatus{
	// 	Name: to.Ptr("D612C429-2526-49D5-961B-885AE11406FD"),
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-01T20:56:36.002Z"); return t}()),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/locations/westus/MediaServicesOperationStatuses/D612C429-2526-49D5-961B-885AE11406FD"),
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-01T20:56:36.002Z"); return t}()),
	// 	Status: to.Ptr("Succeeded"),
	// }
}
