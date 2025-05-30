
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.ApplicationGatewayInner;
import com.azure.resourcemanager.network.fluent.models.ApplicationGatewayIpConfigurationInner;
import com.azure.resourcemanager.network.fluent.models.ApplicationGatewayRequestRoutingRuleInner;
import com.azure.resourcemanager.network.fluent.models.ApplicationGatewaySslCertificateInner;
import com.azure.resourcemanager.network.models.ApplicationGatewayBackendAddress;
import com.azure.resourcemanager.network.models.ApplicationGatewayBackendAddressPool;
import com.azure.resourcemanager.network.models.ApplicationGatewayBackendHttpSettings;
import com.azure.resourcemanager.network.models.ApplicationGatewayClientAuthConfiguration;
import com.azure.resourcemanager.network.models.ApplicationGatewayCookieBasedAffinity;
import com.azure.resourcemanager.network.models.ApplicationGatewayFrontendIpConfiguration;
import com.azure.resourcemanager.network.models.ApplicationGatewayFrontendPort;
import com.azure.resourcemanager.network.models.ApplicationGatewayGlobalConfiguration;
import com.azure.resourcemanager.network.models.ApplicationGatewayHeaderConfiguration;
import com.azure.resourcemanager.network.models.ApplicationGatewayHttpListener;
import com.azure.resourcemanager.network.models.ApplicationGatewayProtocol;
import com.azure.resourcemanager.network.models.ApplicationGatewayRequestRoutingRuleType;
import com.azure.resourcemanager.network.models.ApplicationGatewayRewriteRule;
import com.azure.resourcemanager.network.models.ApplicationGatewayRewriteRuleActionSet;
import com.azure.resourcemanager.network.models.ApplicationGatewayRewriteRuleCondition;
import com.azure.resourcemanager.network.models.ApplicationGatewayRewriteRuleSet;
import com.azure.resourcemanager.network.models.ApplicationGatewaySku;
import com.azure.resourcemanager.network.models.ApplicationGatewaySkuName;
import com.azure.resourcemanager.network.models.ApplicationGatewaySslCipherSuite;
import com.azure.resourcemanager.network.models.ApplicationGatewaySslPolicy;
import com.azure.resourcemanager.network.models.ApplicationGatewaySslPolicyType;
import com.azure.resourcemanager.network.models.ApplicationGatewaySslProfile;
import com.azure.resourcemanager.network.models.ApplicationGatewaySslProtocol;
import com.azure.resourcemanager.network.models.ApplicationGatewayTier;
import com.azure.resourcemanager.network.models.ApplicationGatewayTrustedClientCertificate;
import com.azure.resourcemanager.network.models.ApplicationGatewayTrustedRootCertificate;
import com.azure.resourcemanager.network.models.ApplicationGatewayUrlConfiguration;
import com.azure.resourcemanager.network.models.ManagedServiceIdentity;
import com.azure.resourcemanager.network.models.ManagedServiceIdentityUserAssignedIdentities;
import com.azure.resourcemanager.network.models.ResourceIdentityType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ApplicationGateways CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/ApplicationGatewayCreate.json
     */
    /**
     * Sample code: Create Application Gateway.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createApplicationGateway(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getApplicationGateways().createOrUpdate("rg1", "appgw",
            new ApplicationGatewayInner().withLocation("eastus").withIdentity(new ManagedServiceIdentity()
                .withType(ResourceIdentityType.USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1",
                    new ManagedServiceIdentityUserAssignedIdentities())))
                .withSku(new ApplicationGatewaySku().withName(ApplicationGatewaySkuName.STANDARD_V2)
                    .withTier(ApplicationGatewayTier.STANDARD_V2).withCapacity(3))
                .withGatewayIpConfigurations(Arrays.asList(new ApplicationGatewayIpConfigurationInner()
                    .withName("appgwipc")
                    .withSubnet(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet/subnets/appgwsubnet"))))
                .withTrustedRootCertificates(
                    Arrays.asList(new ApplicationGatewayTrustedRootCertificate().withName("rootcert").withData("****"),
                        new ApplicationGatewayTrustedRootCertificate().withName("rootcert1")
                            .withKeyVaultSecretId("fakeTokenPlaceholder")))
                .withTrustedClientCertificates(Arrays
                    .asList(new ApplicationGatewayTrustedClientCertificate().withName("clientcert").withData("****")))
                .withSslCertificates(Arrays.asList(
                    new ApplicationGatewaySslCertificateInner().withName("sslcert").withData("****")
                        .withPassword("fakeTokenPlaceholder"),
                    new ApplicationGatewaySslCertificateInner().withName("sslcert2")
                        .withKeyVaultSecretId("fakeTokenPlaceholder")))
                .withFrontendIpConfigurations(Arrays.asList(new ApplicationGatewayFrontendIpConfiguration()
                    .withName("appgwfip")
                    .withPublicIpAddress(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/appgwpip"))))
                .withFrontendPorts(Arrays.asList(new ApplicationGatewayFrontendPort().withName("appgwfp").withPort(443),
                    new ApplicationGatewayFrontendPort().withName("appgwfp80").withPort(80)))
                .withBackendAddressPools(
                    Arrays.asList(
                        new ApplicationGatewayBackendAddressPool().withName("appgwpool")
                            .withBackendAddresses(Arrays.asList(new ApplicationGatewayBackendAddress()
                                .withIpAddress("10.0.1.1"),
                                new ApplicationGatewayBackendAddress().withIpAddress("10.0.1.2"))),
                        new ApplicationGatewayBackendAddressPool().withId(
                            "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendAddressPools/appgwpool1")
                            .withName("appgwpool1").withBackendAddresses(
                                Arrays.asList(new ApplicationGatewayBackendAddress().withIpAddress("10.0.0.1"),
                                    new ApplicationGatewayBackendAddress().withIpAddress("10.0.0.2")))))
                .withBackendHttpSettingsCollection(Arrays.asList(new ApplicationGatewayBackendHttpSettings()
                    .withName("appgwbhs").withPort(80).withProtocol(ApplicationGatewayProtocol.HTTP)
                    .withCookieBasedAffinity(ApplicationGatewayCookieBasedAffinity.DISABLED).withRequestTimeout(30)))
                .withHttpListeners(Arrays.asList(new ApplicationGatewayHttpListener().withName("appgwhl")
                    .withFrontendIpConfiguration(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendIPConfigurations/appgwfip"))
                    .withFrontendPort(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendPorts/appgwfp"))
                    .withProtocol(ApplicationGatewayProtocol.HTTPS)
                    .withSslCertificate(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/sslCertificates/sslcert"))
                    .withSslProfile(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/sslProfiles/sslProfile1"))
                    .withRequireServerNameIndication(false),
                    new ApplicationGatewayHttpListener().withName("appgwhttplistener")
                        .withFrontendIpConfiguration(new SubResource().withId(
                            "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendIPConfigurations/appgwfip"))
                        .withFrontendPort(new SubResource().withId(
                            "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendPorts/appgwfp80"))
                        .withProtocol(ApplicationGatewayProtocol.HTTP)))
                .withSslProfiles(Arrays.asList(new ApplicationGatewaySslProfile()
                    .withName("sslProfile1")
                    .withTrustedClientCertificates(Arrays.asList(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/trustedClientCertificates/clientcert")))
                    .withSslPolicy(new ApplicationGatewaySslPolicy()
                        .withPolicyType(ApplicationGatewaySslPolicyType.CUSTOM)
                        .withCipherSuites(
                            Arrays.asList(ApplicationGatewaySslCipherSuite.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256))
                        .withMinProtocolVersion(ApplicationGatewaySslProtocol.TLSV1_1))
                    .withClientAuthConfiguration(
                        new ApplicationGatewayClientAuthConfiguration().withVerifyClientCertIssuerDN(true))))
                .withRequestRoutingRules(Arrays.asList(new ApplicationGatewayRequestRoutingRuleInner()
                    .withName("appgwrule").withRuleType(ApplicationGatewayRequestRoutingRuleType.BASIC).withPriority(10)
                    .withBackendAddressPool(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendAddressPools/appgwpool"))
                    .withBackendHttpSettings(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendHttpSettingsCollection/appgwbhs"))
                    .withHttpListener(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/httpListeners/appgwhl"))
                    .withRewriteRuleSet(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/rewriteRuleSets/rewriteRuleSet1"))))
                .withRewriteRuleSets(
                    Arrays
                        .asList(new ApplicationGatewayRewriteRuleSet().withName("rewriteRuleSet1")
                            .withRewriteRules(Arrays.asList(new ApplicationGatewayRewriteRule()
                                .withName("Set X-Forwarded-For").withRuleSequence(102)
                                .withConditions(Arrays.asList(new ApplicationGatewayRewriteRuleCondition()
                                    .withVariable("http_req_Authorization").withPattern("^Bearer").withIgnoreCase(true)
                                    .withNegate(false)))
                                .withActionSet(new ApplicationGatewayRewriteRuleActionSet()
                                    .withRequestHeaderConfigurations(Arrays.asList(
                                        new ApplicationGatewayHeaderConfiguration().withHeaderName("X-Forwarded-For")
                                            .withHeaderValue("{var_add_x_forwarded_for_proxy}")))
                                    .withResponseHeaderConfigurations(
                                        Arrays.asList(new ApplicationGatewayHeaderConfiguration()
                                            .withHeaderName("Strict-Transport-Security")
                                            .withHeaderValue("max-age=31536000")))
                                    .withUrlConfiguration(
                                        new ApplicationGatewayUrlConfiguration().withModifiedPath("/abc")))))))
                .withGlobalConfiguration(new ApplicationGatewayGlobalConfiguration().withEnableRequestBuffering(true)
                    .withEnableResponseBuffering(true)),
            com.azure.core.util.Context.NONE);
    }

    // Use "Map.of" if available
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
