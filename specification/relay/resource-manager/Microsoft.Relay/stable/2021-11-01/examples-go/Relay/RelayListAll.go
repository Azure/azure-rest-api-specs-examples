package armrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/Relay/RelayListAll.json
func ExampleWCFRelaysClient_NewListByNamespacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWCFRelaysClient().NewListByNamespacePager("resourcegroup", "example-RelayNamespace-01", nil)
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
		// page.WcfRelaysListResult = armrelay.WcfRelaysListResult{
		// 	Value: []*armrelay.WcfRelay{
		// 		{
		// 			Name: to.Ptr("example-Relay-Wcf-01"),
		// 			Type: to.Ptr("Microsoft.Relay/Namespaces/WcfRelays"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/RG1-eg/providers/Microsoft.Relay/namespaces/example-RelayNamespace-01/WcfRelays/example-Relay-Wcf-01"),
		// 			Properties: &armrelay.WcfRelayProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-24T00:46:27.004Z"); return t}()),
		// 				IsDynamic: to.Ptr(false),
		// 				RelayType: to.Ptr(armrelay.RelaytypeNetTCP),
		// 				RequiresClientAuthorization: to.Ptr(true),
		// 				RequiresTransportSecurity: to.Ptr(true),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-24T00:46:27.004Z"); return t}()),
		// 				UserMetadata: to.Ptr("usermetadata is a placeholder to store user-defined string data for the HybridConnection endpoint.e.g. it can be used to store  descriptive data, such as list of teams and their contact information also user-defined configuration settings can be stored"),
		// 			},
		// 	}},
		// }
	}
}
