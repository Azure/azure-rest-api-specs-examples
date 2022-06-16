package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateGatewayHostnameConfiguration.json
func ExampleGatewayHostnameConfigurationClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapimanagement.NewGatewayHostnameConfigurationClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"rg1",
		"apimService1",
		"gw1",
		"default",
		armapimanagement.GatewayHostnameConfigurationContract{
			Properties: &armapimanagement.GatewayHostnameConfigurationContractProperties{
				CertificateID:              to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/certificates/cert1"),
				Hostname:                   to.Ptr("*"),
				HTTP2Enabled:               to.Ptr(true),
				NegotiateClientCertificate: to.Ptr(false),
				Tls10Enabled:               to.Ptr(false),
				Tls11Enabled:               to.Ptr(false),
			},
		},
		&armapimanagement.GatewayHostnameConfigurationClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
