Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.eventhubs.fluent.models.NetworkRuleSetInner;
import com.azure.resourcemanager.eventhubs.models.DefaultAction;
import com.azure.resourcemanager.eventhubs.models.NWRuleSetIpRules;
import com.azure.resourcemanager.eventhubs.models.NWRuleSetVirtualNetworkRules;
import com.azure.resourcemanager.eventhubs.models.NetworkRuleIpAction;
import com.azure.resourcemanager.eventhubs.models.Subnet;
import java.util.Arrays;

/** Samples for Namespaces CreateOrUpdateNetworkRuleSet. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/NameSpaces/VirtualNetworkRule/EHNetworkRuleSetCreate.json
     */
    /**
     * Sample code: NameSpaceNetworkRuleSetCreate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameSpaceNetworkRuleSetCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .eventHubs()
            .manager()
            .serviceClient()
            .getNamespaces()
            .createOrUpdateNetworkRuleSetWithResponse(
                "ResourceGroup",
                "sdk-Namespace-6019",
                new NetworkRuleSetInner()
                    .withDefaultAction(DefaultAction.DENY)
                    .withVirtualNetworkRules(
                        Arrays
                            .asList(
                                new NWRuleSetVirtualNetworkRules()
                                    .withSubnet(
                                        new Subnet()
                                            .withId(
                                                "/subscriptions/subscriptionid/resourcegroups/resourcegroupid/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet2"))
                                    .withIgnoreMissingVnetServiceEndpoint(true),
                                new NWRuleSetVirtualNetworkRules()
                                    .withSubnet(
                                        new Subnet()
                                            .withId(
                                                "/subscriptions/subscriptionid/resourcegroups/resourcegroupid/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet3"))
                                    .withIgnoreMissingVnetServiceEndpoint(false),
                                new NWRuleSetVirtualNetworkRules()
                                    .withSubnet(
                                        new Subnet()
                                            .withId(
                                                "/subscriptions/subscriptionid/resourcegroups/resourcegroupid/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet6"))
                                    .withIgnoreMissingVnetServiceEndpoint(false)))
                    .withIpRules(
                        Arrays
                            .asList(
                                new NWRuleSetIpRules().withIpMask("1.1.1.1").withAction(NetworkRuleIpAction.ALLOW),
                                new NWRuleSetIpRules().withIpMask("1.1.1.2").withAction(NetworkRuleIpAction.ALLOW),
                                new NWRuleSetIpRules().withIpMask("1.1.1.3").withAction(NetworkRuleIpAction.ALLOW),
                                new NWRuleSetIpRules().withIpMask("1.1.1.4").withAction(NetworkRuleIpAction.ALLOW),
                                new NWRuleSetIpRules().withIpMask("1.1.1.5").withAction(NetworkRuleIpAction.ALLOW))),
                Context.NONE);
    }
}
```
