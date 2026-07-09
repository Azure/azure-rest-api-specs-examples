package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v3"
)

// Generated from example definition: 2026-03-01-preview/WebhookListEvents.json
func ExampleWebhooksClient_NewListEventsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
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
		// page = armcontainerregistry.WebhooksClientListEventsResponse{
		// 	EventListResult: armcontainerregistry.EventListResult{
		// 		Value: []*armcontainerregistry.Event{
		// 			{
		// 				ID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				EventRequestMessage: &armcontainerregistry.EventRequestMessage{
		// 					Content: &armcontainerregistry.EventContent{
		// 						ID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T23:14:37.0707808Z"); return t}()),
		// 						Action: to.Ptr("push"),
		// 						Target: &armcontainerregistry.Target{
		// 							MediaType: to.Ptr("application/vnd.docker.distribution.manifest.v2+json"),
		// 							Size: to.Ptr[int64](708),
		// 							Digest: to.Ptr("sha256:fea8895f450959fa676bcc1df0611ea93823a735a01205fd8622846041d0c7cf"),
		// 							Length: to.Ptr[int64](708),
		// 							Repository: to.Ptr("hello-world"),
		// 							URL: to.Ptr("http://192.168.100.227:5000/v2/hello-world/manifests/sha256:fea8895f450959fa676bcc1df0611ea93823a735a01205fd8622846041d0c7cf"),
		// 							Tag: to.Ptr("latest"),
		// 						},
		// 						Request: &armcontainerregistry.Request{
		// 							ID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 							Addr: to.Ptr("192.168.64.11:42961"),
		// 							Host: to.Ptr("192.168.100.227:5000"),
		// 							Method: to.Ptr("GET"),
		// 							Useragent: to.Ptr("curl/7.38.0"),
		// 						},
		// 						Actor: &armcontainerregistry.Actor{
		// 						},
		// 						Source: &armcontainerregistry.Source{
		// 							Addr: to.Ptr("xtal.local:5000"),
		// 							InstanceID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						},
		// 					},
		// 					Headers: map[string]*string{
		// 						"Content-Type": to.Ptr("application/json"),
		// 						"Content-Length": to.Ptr("719"),
		// 						"Authorization": to.Ptr("******"),
		// 					},
		// 					Method: to.Ptr("POST"),
		// 					RequestURI: to.Ptr("http://myservice.com"),
		// 					Version: to.Ptr("1.1"),
		// 				},
		// 				EventResponseMessage: &armcontainerregistry.EventResponseMessage{
		// 					Headers: map[string]*string{
		// 						"Content-Length": to.Ptr("0"),
		// 					},
		// 					StatusCode: to.Ptr("200"),
		// 					Version: to.Ptr("1.1"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
