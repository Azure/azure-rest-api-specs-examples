import com.azure.core.management.SubResource;
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.FirewallPolicyInner;
import com.azure.resourcemanager.network.models.AzureFirewallThreatIntelMode;
import com.azure.resourcemanager.network.models.DnsSettings;
import com.azure.resourcemanager.network.models.ExplicitProxySettings;
import com.azure.resourcemanager.network.models.FirewallPolicyCertificateAuthority;
import com.azure.resourcemanager.network.models.FirewallPolicyInsights;
import com.azure.resourcemanager.network.models.FirewallPolicyIntrusionDetection;
import com.azure.resourcemanager.network.models.FirewallPolicyIntrusionDetectionBypassTrafficSpecifications;
import com.azure.resourcemanager.network.models.FirewallPolicyIntrusionDetectionConfiguration;
import com.azure.resourcemanager.network.models.FirewallPolicyIntrusionDetectionProtocol;
import com.azure.resourcemanager.network.models.FirewallPolicyIntrusionDetectionSignatureSpecification;
import com.azure.resourcemanager.network.models.FirewallPolicyIntrusionDetectionStateType;
import com.azure.resourcemanager.network.models.FirewallPolicyLogAnalyticsResources;
import com.azure.resourcemanager.network.models.FirewallPolicyLogAnalyticsWorkspace;
import com.azure.resourcemanager.network.models.FirewallPolicySku;
import com.azure.resourcemanager.network.models.FirewallPolicySkuTier;
import com.azure.resourcemanager.network.models.FirewallPolicySnat;
import com.azure.resourcemanager.network.models.FirewallPolicySql;
import com.azure.resourcemanager.network.models.FirewallPolicyThreatIntelWhitelist;
import com.azure.resourcemanager.network.models.FirewallPolicyTransportSecurity;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for FirewallPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/FirewallPolicyPut.json
     */
    /**
     * Sample code: Create FirewallPolicy.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createFirewallPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getFirewallPolicies()
            .createOrUpdate(
                "rg1",
                "firewallPolicy",
                new FirewallPolicyInner()
                    .withLocation("West US")
                    .withTags(mapOf("key1", "value1"))
                    .withThreatIntelMode(AzureFirewallThreatIntelMode.ALERT)
                    .withThreatIntelWhitelist(
                        new FirewallPolicyThreatIntelWhitelist()
                            .withIpAddresses(Arrays.asList("20.3.4.5"))
                            .withFqdns(Arrays.asList("*.microsoft.com")))
                    .withInsights(
                        new FirewallPolicyInsights()
                            .withIsEnabled(true)
                            .withRetentionDays(100)
                            .withLogAnalyticsResources(
                                new FirewallPolicyLogAnalyticsResources()
                                    .withWorkspaces(
                                        Arrays
                                            .asList(
                                                new FirewallPolicyLogAnalyticsWorkspace()
                                                    .withRegion("westus")
                                                    .withWorkspaceId(
                                                        new SubResource()
                                                            .withId(
                                                                "/subscriptions/subid/resourcegroups/rg1/providers/microsoft.operationalinsights/workspaces/workspace1")),
                                                new FirewallPolicyLogAnalyticsWorkspace()
                                                    .withRegion("eastus")
                                                    .withWorkspaceId(
                                                        new SubResource()
                                                            .withId(
                                                                "/subscriptions/subid/resourcegroups/rg1/providers/microsoft.operationalinsights/workspaces/workspace2"))))
                                    .withDefaultWorkspaceId(
                                        new SubResource()
                                            .withId(
                                                "/subscriptions/subid/resourcegroups/rg1/providers/microsoft.operationalinsights/workspaces/defaultWorkspace"))))
                    .withSnat(new FirewallPolicySnat().withPrivateRanges(Arrays.asList("IANAPrivateRanges")))
                    .withSql(new FirewallPolicySql().withAllowSqlRedirect(true))
                    .withDnsSettings(
                        new DnsSettings()
                            .withServers(Arrays.asList("30.3.4.5"))
                            .withEnableProxy(true)
                            .withRequireProxyForNetworkRules(false))
                    .withExplicitProxySettings(
                        new ExplicitProxySettings()
                            .withEnableExplicitProxy(true)
                            .withHttpPort(8087)
                            .withHttpsPort(8087)
                            .withPacFilePort(8087)
                            .withPacFile(
                                "https://tinawstorage.file.core.windows.net/?sv=2020-02-10&ss=bfqt&srt=sco&sp=rwdlacuptfx&se=2021-06-04T07:01:12Z&st=2021-06-03T23:01:12Z&sip=68.65.171.11&spr=https&sig=Plsa0RRVpGbY0IETZZOT6znOHcSro71LLTTbzquYPgs%3D"))
                    .withIntrusionDetection(
                        new FirewallPolicyIntrusionDetection()
                            .withMode(FirewallPolicyIntrusionDetectionStateType.ALERT)
                            .withConfiguration(
                                new FirewallPolicyIntrusionDetectionConfiguration()
                                    .withSignatureOverrides(
                                        Arrays
                                            .asList(
                                                new FirewallPolicyIntrusionDetectionSignatureSpecification()
                                                    .withId("2525004")
                                                    .withMode(FirewallPolicyIntrusionDetectionStateType.DENY)))
                                    .withBypassTrafficSettings(
                                        Arrays
                                            .asList(
                                                new FirewallPolicyIntrusionDetectionBypassTrafficSpecifications()
                                                    .withName("bypassRule1")
                                                    .withDescription("Rule 1")
                                                    .withProtocol(FirewallPolicyIntrusionDetectionProtocol.TCP)
                                                    .withSourceAddresses(Arrays.asList("1.2.3.4"))
                                                    .withDestinationAddresses(Arrays.asList("5.6.7.8"))
                                                    .withDestinationPorts(Arrays.asList("*"))))))
                    .withTransportSecurity(
                        new FirewallPolicyTransportSecurity()
                            .withCertificateAuthority(
                                new FirewallPolicyCertificateAuthority()
                                    .withKeyVaultSecretId("https://kv/secret")
                                    .withName("clientcert")))
                    .withSku(new FirewallPolicySku().withTier(FirewallPolicySkuTier.PREMIUM)),
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
