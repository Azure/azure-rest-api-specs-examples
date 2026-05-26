package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementCreateGatewayHostnameConfiguration.json
func ExampleGatewayHostnameConfigurationClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGatewayHostnameConfigurationClient().CreateOrUpdate(ctx, "rg1", "apimService1", "gw1", "default", armapimanagement.GatewayHostnameConfigurationContract{
		Properties: &armapimanagement.GatewayHostnameConfigurationContractProperties{
			CertificateID:              to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/certificates/cert1"),
			Hostname:                   to.Ptr("*"),
			HTTP2Enabled:               to.Ptr(true),
			NegotiateClientCertificate: to.Ptr(false),
			Tls10Enabled:               to.Ptr(false),
			Tls11Enabled:               to.Ptr(false),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapimanagement.GatewayHostnameConfigurationClientCreateOrUpdateResponse{
	// 	GatewayHostnameConfigurationContract: armapimanagement.GatewayHostnameConfigurationContract{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/service/gateways/hostnameConfigurations"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/gateways/gw1/hostnameConfigurations/default"),
	// 		Properties: &armapimanagement.GatewayHostnameConfigurationContractProperties{
	// 			CertificateID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/certificates/cert1"),
	// 			Hostname: to.Ptr("*"),
	// 			HTTP2Enabled: to.Ptr(true),
	// 			NegotiateClientCertificate: to.Ptr(false),
	// 			Tls10Enabled: to.Ptr(false),
	// 			Tls11Enabled: to.Ptr(false),
	// 		},
	// 	},
	// }
}
