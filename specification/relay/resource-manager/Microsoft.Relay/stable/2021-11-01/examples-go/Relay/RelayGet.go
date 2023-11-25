package armrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/Relay/RelayGet.json
func ExampleWCFRelaysClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWCFRelaysClient().Get(ctx, "resourcegroup", "example-RelayNamespace-9953", "example-Relay-Wcf-1194", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WcfRelay = armrelay.WcfRelay{
	// 	Name: to.Ptr("example-Relay-Wcf-1194"),
	// 	Type: to.Ptr("Microsoft.Relay/WcfRelays"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/Default-ServiceBus-WestUS/providers/Microsoft.Relay/namespaces/example-RelayNamespace-9953/WcfRelays/example-Relay-Wcf-1194"),
	// 	Properties: &armrelay.WcfRelayProperties{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-16T00:26:17.501Z"); return t}()),
	// 		IsDynamic: to.Ptr(false),
	// 		ListenerCount: to.Ptr[int32](0),
	// 		RelayType: to.Ptr(armrelay.RelaytypeNetTCP),
	// 		RequiresClientAuthorization: to.Ptr(true),
	// 		RequiresTransportSecurity: to.Ptr(true),
	// 		UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-16T00:26:17.501Z"); return t}()),
	// 	},
	// }
}
