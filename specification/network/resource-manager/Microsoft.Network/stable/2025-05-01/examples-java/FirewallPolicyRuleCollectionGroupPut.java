
import com.azure.resourcemanager.network.fluent.models.FirewallPolicyRuleCollectionGroupInner;
import com.azure.resourcemanager.network.models.FirewallPolicyFilterRuleCollection;
import com.azure.resourcemanager.network.models.FirewallPolicyFilterRuleCollectionAction;
import com.azure.resourcemanager.network.models.FirewallPolicyFilterRuleCollectionActionType;
import com.azure.resourcemanager.network.models.FirewallPolicyRuleNetworkProtocol;
import com.azure.resourcemanager.network.models.NetworkRule;
import java.util.Arrays;

/**
 * Samples for FirewallPolicyRuleCollectionGroups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * FirewallPolicyRuleCollectionGroupPut.json
     */
    /**
     * Sample code: Create Firewall Policy Rule Collection Group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createFirewallPolicyRuleCollectionGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getFirewallPolicyRuleCollectionGroups().createOrUpdate("rg1",
            "firewallPolicy", "ruleCollectionGroup1",
            new FirewallPolicyRuleCollectionGroupInner().withPriority(100)
                .withRuleCollections(Arrays.asList(new FirewallPolicyFilterRuleCollection()
                    .withName("Example-Filter-Rule-Collection").withPriority(100)
                    .withAction(new FirewallPolicyFilterRuleCollectionAction()
                        .withType(FirewallPolicyFilterRuleCollectionActionType.DENY))
                    .withRules(Arrays.asList(new NetworkRule().withName("network-rule1")
                        .withIpProtocols(Arrays.asList(FirewallPolicyRuleNetworkProtocol.TCP))
                        .withSourceAddresses(Arrays.asList("10.1.25.0/24")).withDestinationAddresses(Arrays.asList("*"))
                        .withDestinationPorts(Arrays.asList("*")))))),
            com.azure.core.util.Context.NONE);
    }
}
