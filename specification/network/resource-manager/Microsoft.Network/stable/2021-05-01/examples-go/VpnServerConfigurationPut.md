```go
package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VpnServerConfigurationPut.json
func ExampleVPNServerConfigurationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewVPNServerConfigurationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<vpn-server-configuration-name>",
		armnetwork.VPNServerConfiguration{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"key1": to.Ptr("value1"),
			},
			Properties: &armnetwork.VPNServerConfigurationProperties{
				RadiusClientRootCertificates: []*armnetwork.VPNServerConfigRadiusClientRootCertificate{
					{
						Name:       to.Ptr("<name>"),
						Thumbprint: to.Ptr("<thumbprint>"),
					}},
				RadiusServerRootCertificates: []*armnetwork.VPNServerConfigRadiusServerRootCertificate{
					{
						Name:           to.Ptr("<name>"),
						PublicCertData: to.Ptr("<public-cert-data>"),
					}},
				RadiusServers: []*armnetwork.RadiusServer{
					{
						RadiusServerAddress: to.Ptr("<radius-server-address>"),
						RadiusServerScore:   to.Ptr[int64](25),
						RadiusServerSecret:  to.Ptr("<radius-server-secret>"),
					}},
				VPNClientIPSecPolicies: []*armnetwork.IPSecPolicy{
					{
						DhGroup:             to.Ptr(armnetwork.DhGroupDHGroup14),
						IkeEncryption:       to.Ptr(armnetwork.IkeEncryptionAES256),
						IkeIntegrity:        to.Ptr(armnetwork.IkeIntegritySHA384),
						IPSecEncryption:     to.Ptr(armnetwork.IPSecEncryptionAES256),
						IPSecIntegrity:      to.Ptr(armnetwork.IPSecIntegritySHA256),
						PfsGroup:            to.Ptr(armnetwork.PfsGroupPFS14),
						SaDataSizeKilobytes: to.Ptr[int32](429497),
						SaLifeTimeSeconds:   to.Ptr[int32](86472),
					}},
				VPNClientRevokedCertificates: []*armnetwork.VPNServerConfigVPNClientRevokedCertificate{
					{
						Name:       to.Ptr("<name>"),
						Thumbprint: to.Ptr("<thumbprint>"),
					}},
				VPNClientRootCertificates: []*armnetwork.VPNServerConfigVPNClientRootCertificate{
					{
						Name:           to.Ptr("<name>"),
						PublicCertData: to.Ptr("<public-cert-data>"),
					}},
				VPNProtocols: []*armnetwork.VPNGatewayTunnelingProtocol{
					to.Ptr(armnetwork.VPNGatewayTunnelingProtocolIkeV2)},
			},
		},
		&armnetwork.VPNServerConfigurationsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.5.0/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.
