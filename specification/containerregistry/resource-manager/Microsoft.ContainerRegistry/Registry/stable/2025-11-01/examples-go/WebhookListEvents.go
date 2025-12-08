package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/523ccabf440d8cf1c5b0ea18a8ad1ffedf4902ac/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/Registry/stable/2025-11-01/examples/WebhookListEvents.json
func ExampleWebhooksClient_NewListEventsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWebhooksClient().NewListEventsPager("myResourceGroup", "myRegistry", "myWebhook", nil)
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
		// page.EventListResult = armcontainerregistry.EventListResult{
		// 	Value: []*armcontainerregistry.Event{
		// 		{
		// 			ID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			EventRequestMessage: &armcontainerregistry.EventRequestMessage{
		// 				Method: to.Ptr("POST"),
		// 				Content: &armcontainerregistry.EventContent{
		// 					Action: to.Ptr("push"),
		// 					Actor: &armcontainerregistry.Actor{
		// 					},
		// 					ID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					Source: &armcontainerregistry.Source{
		// 						Addr: to.Ptr("xtal.local:5000"),
		// 						InstanceID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					},
		// 					Target: &armcontainerregistry.Target{
		// 						Digest: to.Ptr("sha256:fea8895f450959fa676bcc1df0611ea93823a735a01205fd8622846041d0c7cf"),
		// 						Length: to.Ptr[int64](708),
		// 						MediaType: to.Ptr("application/vnd.docker.distribution.manifest.v2+json"),
		// 						Repository: to.Ptr("hello-world"),
		// 						Size: to.Ptr[int64](708),
		// 						Tag: to.Ptr("latest"),
		// 						URL: to.Ptr("http://192.168.100.227:5000/v2/hello-world/manifests/sha256:fea8895f450959fa676bcc1df0611ea93823a735a01205fd8622846041d0c7cf"),
		// 					},
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T23:14:37.070Z"); return t}()),
		// 					Request: &armcontainerregistry.Request{
		// 						Method: to.Ptr("GET"),
		// 						Addr: to.Ptr("192.168.64.11:42961"),
		// 						Host: to.Ptr("192.168.100.227:5000"),
		// 						ID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						Useragent: to.Ptr("curl/7.38.0"),
		// 					},
		// 				},
		// 				Headers: map[string]*string{
		// 					"Authorization": to.Ptr("******"),
		// 					"Content-Length": to.Ptr("719"),
		// 					"Content-Type": to.Ptr("application/json"),
		// 				},
		// 				RequestURI: to.Ptr("http://myservice.com"),
		// 				Version: to.Ptr("1.1"),
		// 			},
		// 			EventResponseMessage: &armcontainerregistry.EventResponseMessage{
		// 				Headers: map[string]*string{
		// 					"Content-Length": to.Ptr("0"),
		// 				},
		// 				StatusCode: to.Ptr("200"),
		// 				Version: to.Ptr("1.1"),
		// 			},
		// 	}},
		// }
	}
}
