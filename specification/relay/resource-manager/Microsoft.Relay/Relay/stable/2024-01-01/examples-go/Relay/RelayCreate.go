package armrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay/v2"
)

// Generated from example definition: 2024-01-01/Relay/RelayCreate.json
func ExampleWCFRelaysClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWCFRelaysClient().CreateOrUpdate(ctx, "resourcegroup", "example-RelayNamespace-9953", "example-Relay-Wcf-1194", armrelay.WcfRelay{
		Properties: &armrelay.WcfRelayProperties{
			RelayType:                   to.Ptr(armrelay.RelaytypeNetTCP),
			RequiresClientAuthorization: to.Ptr(true),
			RequiresTransportSecurity:   to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrelay.WCFRelaysClientCreateOrUpdateResponse{
	// 	WcfRelay: &armrelay.WcfRelay{
	// 		Name: to.Ptr("example-Relay-Wcf-1194"),
	// 		Type: to.Ptr("Microsoft.Relay/WcfRelays"),
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Relay/namespaces/example-RelayNamespace-9953/WcfRelays/example-Relay-Wcf-1194"),
	// 		Properties: &armrelay.WcfRelayProperties{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-16T00:26:17.5014661Z"); return t}()),
	// 			IsDynamic: to.Ptr(false),
	// 			RelayType: to.Ptr(armrelay.RelaytypeNetTCP),
	// 			RequiresClientAuthorization: to.Ptr(true),
	// 			RequiresTransportSecurity: to.Ptr(true),
	// 			UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-16T00:26:17.5014661Z"); return t}()),
	// 		},
	// 	},
	// }
}
