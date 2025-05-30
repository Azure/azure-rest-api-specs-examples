
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.AzureFirewallInner;
import com.azure.resourcemanager.network.models.AzureFirewallApplicationRule;
import com.azure.resourcemanager.network.models.AzureFirewallApplicationRuleCollection;
import com.azure.resourcemanager.network.models.AzureFirewallApplicationRuleProtocol;
import com.azure.resourcemanager.network.models.AzureFirewallApplicationRuleProtocolType;
import com.azure.resourcemanager.network.models.AzureFirewallIpConfiguration;
import com.azure.resourcemanager.network.models.AzureFirewallNatRCAction;
import com.azure.resourcemanager.network.models.AzureFirewallNatRCActionType;
import com.azure.resourcemanager.network.models.AzureFirewallNatRule;
import com.azure.resourcemanager.network.models.AzureFirewallNatRuleCollection;
import com.azure.resourcemanager.network.models.AzureFirewallNetworkRule;
import com.azure.resourcemanager.network.models.AzureFirewallNetworkRuleCollection;
import com.azure.resourcemanager.network.models.AzureFirewallNetworkRuleProtocol;
import com.azure.resourcemanager.network.models.AzureFirewallRCAction;
import com.azure.resourcemanager.network.models.AzureFirewallRCActionType;
import com.azure.resourcemanager.network.models.AzureFirewallSku;
import com.azure.resourcemanager.network.models.AzureFirewallSkuName;
import com.azure.resourcemanager.network.models.AzureFirewallSkuTier;
import com.azure.resourcemanager.network.models.AzureFirewallThreatIntelMode;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AzureFirewalls CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/AzureFirewallPutWithZones.
     * json
     */
    /**
     * Sample code: Create Azure Firewall With Zones.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAzureFirewallWithZones(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getAzureFirewalls().createOrUpdate("rg1", "azurefirewall",
            new AzureFirewallInner().withLocation("West US 2").withTags(mapOf("key1", "fakeTokenPlaceholder"))
                .withZones(Arrays.asList("1", "2", "3"))
                .withApplicationRuleCollections(Arrays.asList(new AzureFirewallApplicationRuleCollection().withId(
                    "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall/applicationRuleCollections/apprulecoll")
                    .withName("apprulecoll").withPriority(110)
                    .withAction(new AzureFirewallRCAction().withType(AzureFirewallRCActionType.DENY))
                    .withRules(Arrays.asList(
                        new AzureFirewallApplicationRule().withName("rule1").withDescription("Deny inbound rule")
                            .withSourceAddresses(Arrays.asList("216.58.216.164", "10.0.0.0/24"))
                            .withProtocols(Arrays.asList(new AzureFirewallApplicationRuleProtocol()
                                .withProtocolType(AzureFirewallApplicationRuleProtocolType.HTTPS).withPort(443)))
                            .withTargetFqdns(Arrays.asList("www.test.com"))))))
                .withNatRuleCollections(Arrays.asList(new AzureFirewallNatRuleCollection().withId(
                    "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall/natRuleCollections/natrulecoll")
                    .withName("natrulecoll").withPriority(112)
                    .withAction(new AzureFirewallNatRCAction().withType(AzureFirewallNatRCActionType.DNAT))
                    .withRules(Arrays.asList(
                        new AzureFirewallNatRule().withName("DNAT-HTTPS-traffic")
                            .withDescription("D-NAT all outbound web traffic for inspection")
                            .withSourceAddresses(Arrays.asList("*")).withDestinationAddresses(Arrays.asList("1.2.3.4"))
                            .withDestinationPorts(Arrays.asList("443"))
                            .withProtocols(Arrays.asList(AzureFirewallNetworkRuleProtocol.TCP))
                            .withTranslatedAddress("1.2.3.5").withTranslatedPort("8443"),
                        new AzureFirewallNatRule().withName("DNAT-HTTP-traffic-With-FQDN")
                            .withDescription("D-NAT all inbound web traffic for inspection")
                            .withSourceAddresses(Arrays.asList("*")).withDestinationAddresses(Arrays.asList("1.2.3.4"))
                            .withDestinationPorts(Arrays.asList("80"))
                            .withProtocols(Arrays.asList(AzureFirewallNetworkRuleProtocol.TCP))
                            .withTranslatedPort("880").withTranslatedFqdn("internalhttpserver")))))
                .withNetworkRuleCollections(Arrays.asList(new AzureFirewallNetworkRuleCollection().withId(
                    "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall/networkRuleCollections/netrulecoll")
                    .withName("netrulecoll").withPriority(112).withAction(
                        new AzureFirewallRCAction().withType(AzureFirewallRCActionType.DENY))
                    .withRules(
                        Arrays
                            .asList(
                                new AzureFirewallNetworkRule().withName("L4-traffic")
                                    .withDescription("Block traffic based on source IPs and ports")
                                    .withProtocols(Arrays.asList(AzureFirewallNetworkRuleProtocol.TCP))
                                    .withSourceAddresses(
                                        Arrays.asList("192.168.1.1-192.168.1.12", "10.1.4.12-10.1.4.255"))
                                    .withDestinationAddresses(Arrays.asList("*"))
                                    .withDestinationPorts(Arrays.asList("443-444", "8443")),
                                new AzureFirewallNetworkRule().withName("L4-traffic-with-FQDN")
                                    .withDescription("Block traffic based on source IPs and ports to amazon")
                                    .withProtocols(Arrays.asList(AzureFirewallNetworkRuleProtocol.TCP))
                                    .withSourceAddresses(Arrays.asList("10.2.4.12-10.2.4.255"))
                                    .withDestinationPorts(Arrays.asList("443-444", "8443"))
                                    .withDestinationFqdns(Arrays.asList("www.amazon.com"))))))
                .withIpConfigurations(Arrays.asList(new AzureFirewallIpConfiguration()
                    .withName("azureFirewallIpConfiguration")
                    .withSubnet(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/AzureFirewallSubnet"))
                    .withPublicIpAddress(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pipName"))))
                .withThreatIntelMode(AzureFirewallThreatIntelMode.ALERT).withSku(new AzureFirewallSku()
                    .withName(AzureFirewallSkuName.AZFW_VNET).withTier(AzureFirewallSkuTier.STANDARD)),
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
