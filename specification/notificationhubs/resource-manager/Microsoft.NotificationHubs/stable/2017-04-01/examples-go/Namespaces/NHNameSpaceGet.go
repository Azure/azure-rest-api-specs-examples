package armnotificationhubs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/Namespaces/NHNameSpaceGet.json
func ExampleNamespacesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnotificationhubs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNamespacesClient().Get(ctx, "5ktrial", "nh-sdk-ns", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NamespaceResource = armnotificationhubs.NamespaceResource{
	// 	Name: to.Ptr("nh-sdk-ns"),
	// 	Type: to.Ptr("Microsoft.NotificationHubs/namespaces"),
	// 	ID: to.Ptr("/subscriptions/29cfa613-cbbc-4512-b1d6-1b3a92c7fa40/resourceGroups/5ktrial/providers/Microsoft.NotificationHubs/namespaces/nh-sdk-ns"),
	// 	Location: to.Ptr("South Central US"),
	// 	SKU: &armnotificationhubs.SKU{
	// 		Name: to.Ptr(armnotificationhubs.SKUNameBasic),
	// 	},
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armnotificationhubs.NamespaceProperties{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-02T00:44:56.580Z"); return t}()),
	// 		Critical: to.Ptr(false),
	// 		DataCenter: to.Ptr("SN1"),
	// 		Enabled: to.Ptr(true),
	// 		NamespaceType: to.Ptr(armnotificationhubs.NamespaceTypeNotificationHub),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ScaleUnit: to.Ptr("SN1-001"),
	// 		ServiceBusEndpoint: to.Ptr("https://nh-sdk-ns.servicebus.windows.net:443/"),
	// 		Status: to.Ptr("Active"),
	// 		UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-02T01:02:19.790Z"); return t}()),
	// 	},
	// }
}
