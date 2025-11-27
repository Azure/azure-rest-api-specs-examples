package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/GetSiteInstanceInfo.json
func ExampleWebAppsClient_GetInstanceInfoSlot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWebAppsClient().GetInstanceInfoSlot(ctx, "testrg123", "tests346", "134987120", "staging", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WebSiteInstanceStatus = armappservice.WebSiteInstanceStatus{
	// 	Name: to.Ptr("134987120"),
	// 	Type: to.Ptr("Microsoft.Web/sites/instances"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/sites/tests346/slot/staging/instances/134987120"),
	// 	Properties: &armappservice.WebSiteInstanceStatusProperties{
	// 		Containers: map[string]*armappservice.ContainerInfo{
	// 			"c1": &armappservice.ContainerInfo{
	// 				CurrentTimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2013-10-21T13:28:06.419Z"); return t}()),
	// 			},
	// 		},
	// 		DetectorURL: to.Ptr("testsec579"),
	// 		State: to.Ptr(armappservice.SiteRuntimeStateREADY),
	// 		StatusURL: to.Ptr("https://sampleurl"),
	// 	},
	// }
}
