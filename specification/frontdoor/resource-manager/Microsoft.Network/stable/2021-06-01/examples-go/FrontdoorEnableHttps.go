package armfrontdoor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/FrontdoorEnableHttps.json
func ExampleFrontendEndpointsClient_BeginEnableHTTPS() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfrontdoor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFrontendEndpointsClient().BeginEnableHTTPS(ctx, "rg1", "frontDoor1", "frontendEndpoint1", armfrontdoor.CustomHTTPSConfiguration{
		CertificateSource: to.Ptr(armfrontdoor.FrontDoorCertificateSourceAzureKeyVault),
		KeyVaultCertificateSourceParameters: &armfrontdoor.KeyVaultCertificateSourceParameters{
			SecretName:    to.Ptr("secret1"),
			SecretVersion: to.Ptr("00000000-0000-0000-0000-000000000000"),
			Vault: &armfrontdoor.KeyVaultCertificateSourceParametersVault{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.KeyVault/vaults/vault1"),
			},
		},
		MinimumTLSVersion: to.Ptr(armfrontdoor.MinimumTLSVersionOne0),
		ProtocolType:      to.Ptr(armfrontdoor.FrontDoorTLSProtocolTypeServerNameIndication),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
