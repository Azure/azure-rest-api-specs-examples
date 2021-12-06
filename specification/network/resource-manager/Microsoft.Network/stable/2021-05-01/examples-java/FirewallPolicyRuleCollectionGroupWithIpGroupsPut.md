Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.FirewallPolicyRuleCollectionGroupInner;
import com.azure.resourcemanager.network.models.FirewallPolicyFilterRuleCollection;
import com.azure.resourcemanager.network.models.FirewallPolicyFilterRuleCollectionAction;
import com.azure.resourcemanager.network.models.FirewallPolicyFilterRuleCollectionActionType;
import com.azure.resourcemanager.network.models.FirewallPolicyRuleNetworkProtocol;
import com.azure.resourcemanager.network.models.NetworkRule;
import java.util.Arrays;

/** Samples for FirewallPolicyRuleCollectionGroups CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/FirewallPolicyRuleCollectionGroupWithIpGroupsPut.json
     */
    /**
     * Sample code: Create FirewallPolicyRuleCollectionGroup With IpGroups.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createFirewallPolicyRuleCollectionGroupWithIpGroups(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getFirewallPolicyRuleCollectionGroups()
            .createOrUpdate(
                "rg1",
                "firewallPolicy",
                "ruleCollectionGroup1",
                new FirewallPolicyRuleCollectionGroupInner()
                    .withPriority(110)
                    .withRuleCollections(
                        Arrays
                            .asList(
                                new FirewallPolicyFilterRuleCollection()
                                    .withName("Example-Filter-Rule-Collection")
                                    .withAction(
                                        new FirewallPolicyFilterRuleCollectionAction()
                                            .withType(FirewallPolicyFilterRuleCollectionActionType.DENY))
                                    .withRules(
                                        Arrays
                                            .asList(
                                                new NetworkRule()
                                                    .withName("network-1")
                                                    .withIpProtocols(
                                                        Arrays.asList(FirewallPolicyRuleNetworkProtocol.TCP))
                                                    .withDestinationPorts(Arrays.asList("*"))
                                                    .withSourceIpGroups(
                                                        Arrays
                                                            .asList(
                                                                "/subscriptions/subid/providers/Microsoft.Network/resourceGroup/rg1/ipGroups/ipGroups1"))
                                                    .withDestinationIpGroups(
                                                        Arrays
                                                            .asList(
                                                                "/subscriptions/subid/providers/Microsoft.Network/resourceGroup/rg1/ipGroups/ipGroups2")))))),
                Context.NONE);
    }
}
```
