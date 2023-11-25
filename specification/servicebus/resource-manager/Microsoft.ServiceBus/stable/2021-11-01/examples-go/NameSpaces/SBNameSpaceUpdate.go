package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/SBNameSpaceUpdate.json
func ExampleNamespacesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNamespacesClient().Update(ctx, "ArunMonocle", "sdk-Namespace-3285", armservicebus.SBNamespaceUpdateParameters{
		Location: to.Ptr("South Central US"),
		Tags: map[string]*string{
			"tag3": to.Ptr("value3"),
			"tag4": to.Ptr("value4"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SBNamespace = armservicebus.SBNamespace{
	// 	Name: to.Ptr("sdk-Namespace-3285"),
	// 	Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
	// 	ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ArunMonocle/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-3285"),
	// 	Location: to.Ptr("South Central US"),
	// 	Tags: map[string]*string{
	// 		"tag3": to.Ptr("value3"),
	// 		"tag4": to.Ptr("value4"),
	// 	},
	// 	Properties: &armservicebus.SBNamespaceProperties{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-25T23:07:58.170Z"); return t}()),
	// 		DisableLocalAuth: to.Ptr(false),
	// 		MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:sdk-namespace-3285"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ServiceBusEndpoint: to.Ptr("https://sdk-Namespace-3285.servicebus.windows-int.net:443/"),
	// 		UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-25T23:08:45.497Z"); return t}()),
	// 	},
	// 	SKU: &armservicebus.SBSKU{
	// 		Name: to.Ptr(armservicebus.SKUNameStandard),
	// 		Tier: to.Ptr(armservicebus.SKUTierStandard),
	// 	},
	// }
}
