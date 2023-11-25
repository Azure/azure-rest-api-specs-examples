package armnotificationhubs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/Namespaces/NHNameSpaceCreate.json
func ExampleNamespacesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnotificationhubs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNamespacesClient().CreateOrUpdate(ctx, "5ktrial", "nh-sdk-ns", armnotificationhubs.NamespaceCreateOrUpdateParameters{
		Location: to.Ptr("South Central US"),
		SKU: &armnotificationhubs.SKU{
			Name: to.Ptr(armnotificationhubs.SKUNameStandard),
			Tier: to.Ptr("Standard"),
		},
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NamespaceResource = armnotificationhubs.NamespaceResource{
	// 	Name: to.Ptr("sdk-Namespace-2924"),
	// 	Type: to.Ptr("Microsoft.NotificationHubs/Namespaces"),
	// 	ID: to.Ptr("/subscriptions/29cfa613-cbbc-4512-b1d6-1b3a92c7fa40/resourceGroups/ArunMonocle/providers/Microsoft.NotificationHubs/namespaces/sdk-Namespace-2924"),
	// 	Location: to.Ptr("South Central US"),
	// 	SKU: &armnotificationhubs.SKU{
	// 		Name: to.Ptr(armnotificationhubs.SKUNameStandard),
	// 		Tier: to.Ptr("Standard"),
	// 	},
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armnotificationhubs.NamespaceProperties{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-25T22:26:36.760Z"); return t}()),
	// 		MetricID: to.Ptr("29cfa613-cbbc-4512-b1d6-1b3a92c7fa40:sdk-namespace-2924"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ServiceBusEndpoint: to.Ptr("https://sdk-Namespace-2924.servicebus.windows-int.net:443/"),
	// 		UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-25T22:26:36.760Z"); return t}()),
	// 	},
	// }
}
