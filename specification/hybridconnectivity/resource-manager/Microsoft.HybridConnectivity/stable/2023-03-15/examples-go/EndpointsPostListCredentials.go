package armhybridconnectivity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/stable/2023-03-15/examples/EndpointsPostListCredentials.json
func ExampleEndpointsClient_ListCredentials() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEndpointsClient().ListCredentials(ctx, "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine", "default", &armhybridconnectivity.EndpointsClientListCredentialsOptions{Expiresin: to.Ptr[int64](10800),
		ListCredentialsRequest: &armhybridconnectivity.ListCredentialsRequest{
			ServiceName: to.Ptr(armhybridconnectivity.ServiceNameSSH),
		},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EndpointAccessResource = armhybridconnectivity.EndpointAccessResource{
	// 	Relay: &armhybridconnectivity.RelayNamespaceAccessProperties{
	// 		AccessKey: to.Ptr("SharedAccessSignature sr=http%3A%2F%2Fazgnrelay-eastus-l1.servicebus.windows.net%2Fmicrosoft.kubernetes%2Fconnectedclusters%2Fa0e1fd7d1d974ddf6b11a952d67679c9f12c006eee16861857a8268da4eb1498%2F1619989456957411072%2F&sig=WxDwPF6AmmODaMHNnBGDSm773UG%2B%2Be"),
	// 		ExpiresOn: to.Ptr[int64](1620000256),
	// 		HybridConnectionName: to.Ptr("microsoft.kubernetes/connectedclusters/a0e1fd7d1d974ddf6b11a952d67679c9f12c006eee16861857a8268da4eb1498/1619989456957411072"),
	// 		NamespaceName: to.Ptr("azgnrelay-eastus-l1"),
	// 		NamespaceNameSuffix: to.Ptr("servicebus.windows.net"),
	// 		ServiceConfigurationToken: to.Ptr("SSHvjqH=pTlKql=RtMGw/-k5VFBxSYHIiq5ZgbGFcLkNrDNz5fDsinCN2zkG"),
	// 	},
	// }
}
