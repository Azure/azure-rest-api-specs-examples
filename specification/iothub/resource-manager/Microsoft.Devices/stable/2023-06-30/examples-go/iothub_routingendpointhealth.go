package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c178f2467f792a322f56174be244135d2c907f/specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/iothub_routingendpointhealth.json
func ExampleResourceClient_NewGetEndpointHealthPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiothub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewResourceClient().NewGetEndpointHealthPager("myResourceGroup", "testHub", nil)
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
		// page.EndpointHealthDataListResult = armiothub.EndpointHealthDataListResult{
		// 	Value: []*armiothub.EndpointHealthData{
		// 		{
		// 			EndpointID: to.Ptr("id1"),
		// 			HealthStatus: to.Ptr(armiothub.EndpointHealthStatusHealthy),
		// 			LastSendAttemptTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2022-09-13T18:04:32.000Z"); return t}()),
		// 			LastSuccessfulSendAttemptTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2022-09-13T18:04:32.000Z"); return t}()),
		// 		},
		// 		{
		// 			EndpointID: to.Ptr("id2"),
		// 			HealthStatus: to.Ptr(armiothub.EndpointHealthStatusUnknown),
		// 		},
		// 		{
		// 			EndpointID: to.Ptr("id3"),
		// 			HealthStatus: to.Ptr(armiothub.EndpointHealthStatusUnhealthy),
		// 			LastKnownError: to.Ptr("NotFound"),
		// 			LastKnownErrorTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2022-09-13T18:04:32.000Z"); return t}()),
		// 			LastSendAttemptTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2022-09-13T18:04:32.000Z"); return t}()),
		// 			LastSuccessfulSendAttemptTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2022-09-13T18:04:32.000Z"); return t}()),
		// 	}},
		// }
	}
}
