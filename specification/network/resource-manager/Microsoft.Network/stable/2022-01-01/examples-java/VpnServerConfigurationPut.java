import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.VpnServerConfigurationInner;
import com.azure.resourcemanager.network.fluent.models.VpnServerConfigurationPolicyGroupInner;
import com.azure.resourcemanager.network.models.DhGroup;
import com.azure.resourcemanager.network.models.IkeEncryption;
import com.azure.resourcemanager.network.models.IkeIntegrity;
import com.azure.resourcemanager.network.models.IpsecEncryption;
import com.azure.resourcemanager.network.models.IpsecIntegrity;
import com.azure.resourcemanager.network.models.IpsecPolicy;
import com.azure.resourcemanager.network.models.PfsGroup;
import com.azure.resourcemanager.network.models.RadiusServer;
import com.azure.resourcemanager.network.models.VpnGatewayTunnelingProtocol;
import com.azure.resourcemanager.network.models.VpnPolicyMemberAttributeType;
import com.azure.resourcemanager.network.models.VpnServerConfigRadiusClientRootCertificate;
import com.azure.resourcemanager.network.models.VpnServerConfigRadiusServerRootCertificate;
import com.azure.resourcemanager.network.models.VpnServerConfigVpnClientRevokedCertificate;
import com.azure.resourcemanager.network.models.VpnServerConfigVpnClientRootCertificate;
import com.azure.resourcemanager.network.models.VpnServerConfigurationPolicyGroupMember;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for VpnServerConfigurations CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VpnServerConfigurationPut.json
     */
    /**
     * Sample code: VpnServerConfigurationCreate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnServerConfigurationCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVpnServerConfigurations()
            .createOrUpdate(
                "rg1",
                "vpnServerConfiguration1",
                new VpnServerConfigurationInner()
                    .withLocation("West US")
                    .withTags(mapOf("key1", "value1"))
                    .withVpnProtocols(Arrays.asList(VpnGatewayTunnelingProtocol.IKE_V2))
                    .withVpnClientRootCertificates(
                        Arrays
                            .asList(
                                new VpnServerConfigVpnClientRootCertificate()
                                    .withName("vpnServerConfigVpnClientRootCert1")
                                    .withPublicCertData(
                                        "MIIC5zCCAc+gAwIBAgIQErQ0Hk4aDJxIA+Q5RagB+jANBgkqhkiG9w0BAQsFADAWMRQwEgYDVQQDDAtQMlNSb290Q2VydDAeFw0xNzEyMTQyMTA3MzhaFw0xODEyMTQyMTI3MzhaMBYxFDASBgNVBAMMC1AyU1Jvb3RDZXJ0MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArP7/NQXmW7cQ/ZR1mv3Y3I29Lt7HTOqzo/1KUOoVH3NItbQIRAQbwKy3UWrOFz4eGNX2GWtNRMdCyWsKeqy9Ltsdfcm1IbKXkl84DFeU/ZacXu4Dl3xX3gV5du4TLZjEowJELyur11Ea2YcjPRQ/FzAF9/hGuboS1HZQEPLx4FdUs9OxCYOtc0MxBCwLfVTTRqarb0Ne+arNYd4kCzIhAke1nOyKAJBda5ZL+VHy3S5S8qGlD46jm8HXugmAkUygS4oIIXOmj/1O9sNAi3LN60zufSzCmP8Rm/iUGX+DHAGGiXxwZOKQLEDaZXKqoHjMPP0XudmSWwOIbyeQVrLhkwIDAQABozEwLzAOBgNVHQ8BAf8EBAMCAgQwHQYDVR0OBBYEFEfeNU2trYxNLF9ONmuJUsT13pKDMA0GCSqGSIb3DQEBCwUAA4IBAQBmM6RJzsGGipxyMhimHKN2xlkejhVsgBoTAhOU0llW9aUSwINJ9zFUGgI8IzUFy1VG776fchHp0LMRmPSIUYk5btEPxbsrPtumPuMH8EQGrS+Rt4pD+78c8H1fEPkq5CmDl/PKu4JoFGv+aFcE+Od0hlILstIF10Qysf++QXDolKfzJa/56bgMeYKFiju73loiRM57ns8ddXpfLl792UVpRkFU62LNns6Y1LKTwapmUF4IvIuAIzd6LZNOQng64LAKXtKnViJ1JQiXwf4CEzhgvAti3/ejpb3U90hsrUcyZi6wBv9bZLcAJRWpz61JNYliM1d1grSwQDKGXNQE4xuN")))
                    .withVpnClientRevokedCertificates(
                        Arrays
                            .asList(
                                new VpnServerConfigVpnClientRevokedCertificate()
                                    .withName("vpnServerConfigVpnClientRevokedCert1")
                                    .withThumbprint("83FFBFC8848B5A5836C94D0112367E16148A286F")))
                    .withRadiusServerRootCertificates(
                        Arrays
                            .asList(
                                new VpnServerConfigRadiusServerRootCertificate()
                                    .withName("vpnServerConfigRadiusServerRootCer1")
                                    .withPublicCertData(
                                        "MIIC5zCCAc+gAwIBAgIQErQ0Hk4aDJxIA+Q5RagB+jANBgkqhkiG9w0BAQsFADAWMRQwEgYDVQQDDAtQMlNSb290Q2VydDAeFw0xNzEyMTQyMTA3MzhaFw0xODEyMTQyMTI3MzhaMBYxFDASBgNVBAMMC1AyU1Jvb3RDZXJ0MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArP7/NQXmW7cQ/ZR1mv3Y3I29Lt7HTOqzo/1KUOoVH3NItbQIRAQbwKy3UWrOFz4eGNX2GWtNRMdCyWsKeqy9Ltsdfcm1IbKXkl84DFeU/ZacXu4Dl3xX3gV5du4TLZjEowJELyur11Ea2YcjPRQ/FzAF9/hGuboS1HZQEPLx4FdUs9OxCYOtc0MxBCwLfVTTRqarb0Ne+arNYd4kCzIhAke1nOyKAJBda5ZL+VHy3S5S8qGlD46jm8HXugmAkUygS4oIIXOmj/1O9sNAi3LN60zufSzCmP8Rm/iUGX+DHAGGiXxwZOKQLEDaZXKqoHjMPP0XudmSWwOIbyeQVrLhkwIDAQABozEwLzAOBgNVHQ8BAf8EBAMCAgQwHQYDVR0OBBYEFEfeNU2trYxNLF9ONmuJUsT13pKDMA0GCSqGSIb3DQEBCwUAA4IBAQBmM6RJzsGGipxyMhimHKN2xlkejhVsgBoTAhOU0llW9aUSwINJ9zFUGgI8IzUFy1VG776fchHp0LMRmPSIUYk5btEPxbsrPtumPuMH8EQGrS+Rt4pD+78c8H1fEPkq5CmDl/PKu4JoFGv+aFcE+Od0hlILstIF10Qysf++QXDolKfzJa/56bgMeYKFiju73loiRM57ns8ddXpfLl792UVpRkFU62LNns6Y1LKTwapmUF4IvIuAIzd6LZNOQng64LAKXtKnViJ1JQiXwf4CEzhgvAti3/ejpb3U90hsrUcyZi6wBv9bZLcAJRWpz61JNYliM1d1grSwQDKGXNQE4xuM")))
                    .withRadiusClientRootCertificates(
                        Arrays
                            .asList(
                                new VpnServerConfigRadiusClientRootCertificate()
                                    .withName("vpnServerConfigRadiusClientRootCert1")
                                    .withThumbprint("83FFBFC8848B5A5836C94D0112367E16148A286F")))
                    .withVpnClientIpsecPolicies(
                        Arrays
                            .asList(
                                new IpsecPolicy()
                                    .withSaLifeTimeSeconds(86472)
                                    .withSaDataSizeKilobytes(429497)
                                    .withIpsecEncryption(IpsecEncryption.AES256)
                                    .withIpsecIntegrity(IpsecIntegrity.SHA256)
                                    .withIkeEncryption(IkeEncryption.AES256)
                                    .withIkeIntegrity(IkeIntegrity.SHA384)
                                    .withDhGroup(DhGroup.DHGROUP14)
                                    .withPfsGroup(PfsGroup.PFS14)))
                    .withRadiusServers(
                        Arrays
                            .asList(
                                new RadiusServer()
                                    .withRadiusServerAddress("10.0.0.0")
                                    .withRadiusServerScore(25L)
                                    .withRadiusServerSecret("radiusServerSecret")))
                    .withConfigurationPolicyGroups(
                        Arrays
                            .asList(
                                new VpnServerConfigurationPolicyGroupInner()
                                    .withId(
                                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfiguration1/vpnServerConfigurationPolicyGroups/policyGroup1")
                                    .withName("policyGroup1")
                                    .withIsDefault(true)
                                    .withPriority(0)
                                    .withPolicyMembers(
                                        Arrays
                                            .asList(
                                                new VpnServerConfigurationPolicyGroupMember()
                                                    .withName("policy1")
                                                    .withAttributeType(
                                                        VpnPolicyMemberAttributeType.RADIUS_AZURE_GROUP_ID)
                                                    .withAttributeValue("6ad1bd08"))),
                                new VpnServerConfigurationPolicyGroupInner()
                                    .withId(
                                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfiguration1/vpnServerConfigurationPolicyGroups/policyGroup2")
                                    .withName("policyGroup2")
                                    .withIsDefault(true)
                                    .withPriority(0)
                                    .withPolicyMembers(
                                        Arrays
                                            .asList(
                                                new VpnServerConfigurationPolicyGroupMember()
                                                    .withName("policy2")
                                                    .withAttributeType(
                                                        VpnPolicyMemberAttributeType.CERTIFICATE_GROUP_ID)
                                                    .withAttributeValue("red.com"))))),
                Context.NONE);
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
