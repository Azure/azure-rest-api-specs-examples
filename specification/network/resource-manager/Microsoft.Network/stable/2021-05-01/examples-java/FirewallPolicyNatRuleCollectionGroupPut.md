Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.FirewallPolicyRuleCollectionGroupInner;
import com.azure.resourcemanager.network.models.FirewallPolicyNatRuleCollection;
import com.azure.resourcemanager.network.models.FirewallPolicyNatRuleCollectionAction;
import com.azure.resourcemanager.network.models.FirewallPolicyNatRuleCollectionActionType;
import com.azure.resourcemanager.network.models.FirewallPolicyRuleNetworkProtocol;
import com.azure.resourcemanager.network.models.NatRule;
import java.util.Arrays;

/** Samples for FirewallPolicyRuleCollectionGroups CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/FirewallPolicyNatRuleCollectionGroupPut.json
     */
    /**
     * Sample code: Create FirewallPolicyNatRuleCollectionGroup.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createFirewallPolicyNatRuleCollectionGroup(
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
                    .withPriority(100)
                    .withRuleCollections(
                        Arrays
                            .asList(
                                new FirewallPolicyNatRuleCollection()
                                    .withName("Example-Nat-Rule-Collection")
                                    .withPriority(100)
                                    .withAction(
                                        new FirewallPolicyNatRuleCollectionAction()
                                            .withType(FirewallPolicyNatRuleCollectionActionType.DNAT))
                                    .withRules(
                                        Arrays
                                            .asList(
                                                new NatRule()
                                                    .withName("nat-rule1")
                                                    .withIpProtocols(
                                                        Arrays
                                                            .asList(
                                                                FirewallPolicyRuleNetworkProtocol.TCP,
                                                                FirewallPolicyRuleNetworkProtocol.UDP))
                                                    .withSourceAddresses(Arrays.asList("2.2.2.2"))
                                                    .withDestinationAddresses(Arrays.asList("152.23.32.23"))
                                                    .withDestinationPorts(Arrays.asList("8080"))
                                                    .withTranslatedPort("8080")
                                                    .withSourceIpGroups(Arrays.asList())
                                                    .withTranslatedFqdn("internalhttp.server.net"))))),
                Context.NONE);
    }
}
```
