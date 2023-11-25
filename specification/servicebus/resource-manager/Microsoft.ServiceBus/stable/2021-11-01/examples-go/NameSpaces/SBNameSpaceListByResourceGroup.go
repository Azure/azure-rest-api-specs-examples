package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/SBNameSpaceListByResourceGroup.json
func ExampleNamespacesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNamespacesClient().NewListByResourceGroupPager("ArunMonocle", nil)
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
		// page.SBNamespaceListResult = armservicebus.SBNamespaceListResult{
		// 	Value: []*armservicebus.SBNamespace{
		// 		{
		// 			Name: to.Ptr("sdk-Namespace-2924"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ArunMonocle/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-2924"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 				"tag1": to.Ptr("value1"),
		// 				"tag2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-25T22:26:36.760Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:sdk-namespace-2924"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://sdk-Namespace-2924.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-25T22:26:59.350Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 	}},
		// }
	}
}
