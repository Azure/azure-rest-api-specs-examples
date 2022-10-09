package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/VpnServerConfigurationPut.json
func ExampleVPNServerConfigurationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewVPNServerConfigurationsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "vpnServerConfiguration1", armnetwork.VPNServerConfiguration{
		Location: to.Ptr("West US"),
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
		Properties: &armnetwork.VPNServerConfigurationProperties{
			ConfigurationPolicyGroups: []*armnetwork.VPNServerConfigurationPolicyGroup{
				{
					ID:   to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfiguration1/vpnServerConfigurationPolicyGroups/policyGroup1"),
					Name: to.Ptr("policyGroup1"),
					Properties: &armnetwork.VPNServerConfigurationPolicyGroupProperties{
						IsDefault: to.Ptr(true),
						PolicyMembers: []*armnetwork.VPNServerConfigurationPolicyGroupMember{
							{
								Name:           to.Ptr("policy1"),
								AttributeType:  to.Ptr(armnetwork.VPNPolicyMemberAttributeTypeRadiusAzureGroupID),
								AttributeValue: to.Ptr("6ad1bd08"),
							}},
						Priority: to.Ptr[int32](0),
					},
				},
				{
					ID:   to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfiguration1/vpnServerConfigurationPolicyGroups/policyGroup2"),
					Name: to.Ptr("policyGroup2"),
					Properties: &armnetwork.VPNServerConfigurationPolicyGroupProperties{
						IsDefault: to.Ptr(true),
						PolicyMembers: []*armnetwork.VPNServerConfigurationPolicyGroupMember{
							{
								Name:           to.Ptr("policy2"),
								AttributeType:  to.Ptr(armnetwork.VPNPolicyMemberAttributeTypeCertificateGroupID),
								AttributeValue: to.Ptr("red.com"),
							}},
						Priority: to.Ptr[int32](0),
					},
				}},
			RadiusClientRootCertificates: []*armnetwork.VPNServerConfigRadiusClientRootCertificate{
				{
					Name:       to.Ptr("vpnServerConfigRadiusClientRootCert1"),
					Thumbprint: to.Ptr("83FFBFC8848B5A5836C94D0112367E16148A286F"),
				}},
			RadiusServerRootCertificates: []*armnetwork.VPNServerConfigRadiusServerRootCertificate{
				{
					Name:           to.Ptr("vpnServerConfigRadiusServerRootCer1"),
					PublicCertData: to.Ptr("MIIC5zCCAc+gAwIBAgIQErQ0Hk4aDJxIA+Q5RagB+jANBgkqhkiG9w0BAQsFADAWMRQwEgYDVQQDDAtQMlNSb290Q2VydDAeFw0xNzEyMTQyMTA3MzhaFw0xODEyMTQyMTI3MzhaMBYxFDASBgNVBAMMC1AyU1Jvb3RDZXJ0MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArP7/NQXmW7cQ/ZR1mv3Y3I29Lt7HTOqzo/1KUOoVH3NItbQIRAQbwKy3UWrOFz4eGNX2GWtNRMdCyWsKeqy9Ltsdfcm1IbKXkl84DFeU/ZacXu4Dl3xX3gV5du4TLZjEowJELyur11Ea2YcjPRQ/FzAF9/hGuboS1HZQEPLx4FdUs9OxCYOtc0MxBCwLfVTTRqarb0Ne+arNYd4kCzIhAke1nOyKAJBda5ZL+VHy3S5S8qGlD46jm8HXugmAkUygS4oIIXOmj/1O9sNAi3LN60zufSzCmP8Rm/iUGX+DHAGGiXxwZOKQLEDaZXKqoHjMPP0XudmSWwOIbyeQVrLhkwIDAQABozEwLzAOBgNVHQ8BAf8EBAMCAgQwHQYDVR0OBBYEFEfeNU2trYxNLF9ONmuJUsT13pKDMA0GCSqGSIb3DQEBCwUAA4IBAQBmM6RJzsGGipxyMhimHKN2xlkejhVsgBoTAhOU0llW9aUSwINJ9zFUGgI8IzUFy1VG776fchHp0LMRmPSIUYk5btEPxbsrPtumPuMH8EQGrS+Rt4pD+78c8H1fEPkq5CmDl/PKu4JoFGv+aFcE+Od0hlILstIF10Qysf++QXDolKfzJa/56bgMeYKFiju73loiRM57ns8ddXpfLl792UVpRkFU62LNns6Y1LKTwapmUF4IvIuAIzd6LZNOQng64LAKXtKnViJ1JQiXwf4CEzhgvAti3/ejpb3U90hsrUcyZi6wBv9bZLcAJRWpz61JNYliM1d1grSwQDKGXNQE4xuM"),
				}},
			RadiusServers: []*armnetwork.RadiusServer{
				{
					RadiusServerAddress: to.Ptr("10.0.0.0"),
					RadiusServerScore:   to.Ptr[int64](25),
					RadiusServerSecret:  to.Ptr("radiusServerSecret"),
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
					Name:       to.Ptr("vpnServerConfigVpnClientRevokedCert1"),
					Thumbprint: to.Ptr("83FFBFC8848B5A5836C94D0112367E16148A286F"),
				}},
			VPNClientRootCertificates: []*armnetwork.VPNServerConfigVPNClientRootCertificate{
				{
					Name:           to.Ptr("vpnServerConfigVpnClientRootCert1"),
					PublicCertData: to.Ptr("MIIC5zCCAc+gAwIBAgIQErQ0Hk4aDJxIA+Q5RagB+jANBgkqhkiG9w0BAQsFADAWMRQwEgYDVQQDDAtQMlNSb290Q2VydDAeFw0xNzEyMTQyMTA3MzhaFw0xODEyMTQyMTI3MzhaMBYxFDASBgNVBAMMC1AyU1Jvb3RDZXJ0MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArP7/NQXmW7cQ/ZR1mv3Y3I29Lt7HTOqzo/1KUOoVH3NItbQIRAQbwKy3UWrOFz4eGNX2GWtNRMdCyWsKeqy9Ltsdfcm1IbKXkl84DFeU/ZacXu4Dl3xX3gV5du4TLZjEowJELyur11Ea2YcjPRQ/FzAF9/hGuboS1HZQEPLx4FdUs9OxCYOtc0MxBCwLfVTTRqarb0Ne+arNYd4kCzIhAke1nOyKAJBda5ZL+VHy3S5S8qGlD46jm8HXugmAkUygS4oIIXOmj/1O9sNAi3LN60zufSzCmP8Rm/iUGX+DHAGGiXxwZOKQLEDaZXKqoHjMPP0XudmSWwOIbyeQVrLhkwIDAQABozEwLzAOBgNVHQ8BAf8EBAMCAgQwHQYDVR0OBBYEFEfeNU2trYxNLF9ONmuJUsT13pKDMA0GCSqGSIb3DQEBCwUAA4IBAQBmM6RJzsGGipxyMhimHKN2xlkejhVsgBoTAhOU0llW9aUSwINJ9zFUGgI8IzUFy1VG776fchHp0LMRmPSIUYk5btEPxbsrPtumPuMH8EQGrS+Rt4pD+78c8H1fEPkq5CmDl/PKu4JoFGv+aFcE+Od0hlILstIF10Qysf++QXDolKfzJa/56bgMeYKFiju73loiRM57ns8ddXpfLl792UVpRkFU62LNns6Y1LKTwapmUF4IvIuAIzd6LZNOQng64LAKXtKnViJ1JQiXwf4CEzhgvAti3/ejpb3U90hsrUcyZi6wBv9bZLcAJRWpz61JNYliM1d1grSwQDKGXNQE4xuN"),
				}},
			VPNProtocols: []*armnetwork.VPNGatewayTunnelingProtocol{
				to.Ptr(armnetwork.VPNGatewayTunnelingProtocolIkeV2)},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
