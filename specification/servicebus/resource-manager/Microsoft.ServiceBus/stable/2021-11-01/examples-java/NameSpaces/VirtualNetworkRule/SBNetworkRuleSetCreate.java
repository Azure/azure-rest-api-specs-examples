
import com.azure.resourcemanager.servicebus.fluent.models.NetworkRuleSetInner;
import com.azure.resourcemanager.servicebus.models.DefaultAction;
import com.azure.resourcemanager.servicebus.models.NWRuleSetIpRules;
import com.azure.resourcemanager.servicebus.models.NWRuleSetVirtualNetworkRules;
import com.azure.resourcemanager.servicebus.models.NetworkRuleIpAction;
import com.azure.resourcemanager.servicebus.models.Subnet;
import java.util.Arrays;

/**
 * Samples for Namespaces CreateOrUpdateNetworkRuleSet.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/
     * VirtualNetworkRule/SBNetworkRuleSetCreate.json
     */
    /**
     * Sample code: NameSpaceNetworkRuleSetCreate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameSpaceNetworkRuleSetCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getNamespaces().createOrUpdateNetworkRuleSetWithResponse(
            "ResourceGroup", "sdk-Namespace-6019",
            new NetworkRuleSetInner().withDefaultAction(DefaultAction.DENY).withVirtualNetworkRules(Arrays.asList(
                new NWRuleSetVirtualNetworkRules().withSubnet(new Subnet().withId(
                    "/subscriptions/854d368f-1828-428f-8f3c-f2affa9b2f7d/resourcegroups/alitest/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet2"))
                    .withIgnoreMissingVnetServiceEndpoint(true),
                new NWRuleSetVirtualNetworkRules().withSubnet(new Subnet().withId(
                    "/subscriptions/854d368f-1828-428f-8f3c-f2affa9b2f7d/resourcegroups/alitest/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet3"))
                    .withIgnoreMissingVnetServiceEndpoint(false),
                new NWRuleSetVirtualNetworkRules().withSubnet(new Subnet().withId(
                    "/subscriptions/854d368f-1828-428f-8f3c-f2affa9b2f7d/resourcegroups/alitest/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet6"))
                    .withIgnoreMissingVnetServiceEndpoint(false)))
                .withIpRules(
                    Arrays.asList(new NWRuleSetIpRules().withIpMask("1.1.1.1").withAction(NetworkRuleIpAction.ALLOW),
                        new NWRuleSetIpRules().withIpMask("1.1.1.2").withAction(NetworkRuleIpAction.ALLOW),
                        new NWRuleSetIpRules().withIpMask("1.1.1.3").withAction(NetworkRuleIpAction.ALLOW),
                        new NWRuleSetIpRules().withIpMask("1.1.1.4").withAction(NetworkRuleIpAction.ALLOW),
                        new NWRuleSetIpRules().withIpMask("1.1.1.5").withAction(NetworkRuleIpAction.ALLOW))),
            com.azure.core.util.Context.NONE);
    }
}
