
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
     * x-ms-original-file: 2025-07-01/FirewallPolicyRuleCollectionGroupWithIpGroupsPut.json
     */
    /**
     * Sample code: Create Firewall Policy Rule Collection Group With IP Groups.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        createFirewallPolicyRuleCollectionGroupWithIPGroups(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFirewallPolicyRuleCollectionGroups().createOrUpdate("rg1", "firewallPolicy",
            "ruleCollectionGroup1",
            new FirewallPolicyRuleCollectionGroupInner().withPriority(110)
                .withRuleCollections(Arrays.asList(new FirewallPolicyFilterRuleCollection()
                    .withName("Example-Filter-Rule-Collection")
                    .withAction(new FirewallPolicyFilterRuleCollectionAction()
                        .withType(FirewallPolicyFilterRuleCollectionActionType.DENY))
                    .withRules(Arrays.asList(new NetworkRule().withName("network-1")
                        .withIpProtocols(Arrays.asList(FirewallPolicyRuleNetworkProtocol.TCP))
                        .withDestinationPorts(Arrays.asList("*"))
                        .withSourceIpGroups(Arrays.asList(
                            "/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/resourceGroups/rg1/ipGroups/ipGroups1"))
                        .withDestinationIpGroups(Arrays.asList(
                            "/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/resourceGroups/rg1/ipGroups/ipGroups2")))))),
            com.azure.core.util.Context.NONE);
    }
}
