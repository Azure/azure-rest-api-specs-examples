package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/NameSpaces/SBNameSpaceCreate.json
func ExampleNamespacesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNamespacesClient().BeginCreateOrUpdate(ctx, "ArunMonocle", "sdk-Namespace2924", armservicebus.SBNamespace{
		Location: to.Ptr("South Central US"),
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
		Properties: &armservicebus.SBNamespaceProperties{
			PremiumMessagingPartitions: to.Ptr[int32](2),
		},
		SKU: &armservicebus.SBSKU{
			Name:     to.Ptr(armservicebus.SKUNamePremium),
			Capacity: to.Ptr[int32](4),
			Tier:     to.Ptr(armservicebus.SKUTierPremium),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SBNamespace = armservicebus.SBNamespace{
	// 	Name: to.Ptr("sdk-Namespace-2924"),
	// 	Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
	// 	ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ArunMonocle/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-2924"),
	// 	Location: to.Ptr("South Central US"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armservicebus.SBNamespaceProperties{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-25T22:26:36.76Z"); return t}()),
	// 		MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:sdk-namespace-2924"),
	// 		MinimumTLSVersion: to.Ptr(armservicebus.TLSVersionOne2),
	// 		PremiumMessagingPartitions: to.Ptr[int32](2),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ServiceBusEndpoint: to.Ptr("https://sdk-Namespace-2924.servicebus.windows-int.net:443/"),
	// 		UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-25T22:26:36.76Z"); return t}()),
	// 	},
	// 	SKU: &armservicebus.SBSKU{
	// 		Name: to.Ptr(armservicebus.SKUNamePremium),
	// 		Capacity: to.Ptr[int32](4),
	// 		Tier: to.Ptr(armservicebus.SKUTierPremium),
	// 	},
	// }
}
