package armhybridconnectivity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/stable/2023-03-15/examples/EndpointsPostListIngressGatewayCredentials.json
func ExampleEndpointsClient_ListIngressGatewayCredentials() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEndpointsClient().ListIngressGatewayCredentials(ctx, "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/arcGroup/providers/Microsoft.ArcPlaceHolder/ProvisionedClusters/cluster0", "default", &armhybridconnectivity.EndpointsClientListIngressGatewayCredentialsOptions{Expiresin: to.Ptr[int64](10800),
		ListIngressGatewayCredentialsRequest: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IngressGatewayResource = armhybridconnectivity.IngressGatewayResource{
	// 	Ingress: &armhybridconnectivity.IngressProfileProperties{
	// 		AADProfile: &armhybridconnectivity.AADProfileProperties{
	// 			ServerID: to.Ptr("6256c85f-0aad-4d50-b960-e6e9b21efe35"),
	// 			TenantID: to.Ptr("hy657tgh-8d16-42db-81b7-1234hygt67hy5"),
	// 		},
	// 		Hostname: to.Ptr("clusterhostname"),
	// 	},
	// 	Relay: &armhybridconnectivity.RelayNamespaceAccessProperties{
	// 		AccessKey: to.Ptr("SharedAccessSignature sr=http%3A%2F%2Fazgnrelay-eastus-l1.servicebus.windows.net%2Fmicrosoft.provisionedcluster%hci"),
	// 		ExpiresOn: to.Ptr[int64](1620000256),
	// 		HybridConnectionName: to.Ptr("microsoft.arcplaceholder/provisionedclusters/000/1619989456957411072"),
	// 		NamespaceName: to.Ptr("relaynamespace"),
	// 		NamespaceNameSuffix: to.Ptr("servicebus.windows.net"),
	// 		ServiceConfigurationToken: to.Ptr("SSHvjqH=pTlKql=RtMGw/-k5VFBxSYHIiq5ZgbGFcLkNrDNz5fDsinCN2zkG"),
	// 	},
	// }
}
