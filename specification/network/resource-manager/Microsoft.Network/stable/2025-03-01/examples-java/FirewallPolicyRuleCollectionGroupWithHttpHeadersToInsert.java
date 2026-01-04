
import com.azure.resourcemanager.network.fluent.models.FirewallPolicyRuleCollectionGroupInner;
import com.azure.resourcemanager.network.models.ApplicationRule;
import com.azure.resourcemanager.network.models.FirewallPolicyFilterRuleCollection;
import com.azure.resourcemanager.network.models.FirewallPolicyFilterRuleCollectionAction;
import com.azure.resourcemanager.network.models.FirewallPolicyFilterRuleCollectionActionType;
import com.azure.resourcemanager.network.models.FirewallPolicyHttpHeaderToInsert;
import com.azure.resourcemanager.network.models.FirewallPolicyRuleApplicationProtocol;
import com.azure.resourcemanager.network.models.FirewallPolicyRuleApplicationProtocolType;
import java.util.Arrays;

/**
 * Samples for FirewallPolicyRuleCollectionGroups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * FirewallPolicyRuleCollectionGroupWithHttpHeadersToInsert.json
     */
    /**
     * Sample code: Create Firewall Policy Rule Collection Group With http header to insert.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createFirewallPolicyRuleCollectionGroupWithHttpHeaderToInsert(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getFirewallPolicyRuleCollectionGroups()
            .createOrUpdate("rg1", "firewallPolicy", "ruleCollectionGroup1",
                new FirewallPolicyRuleCollectionGroupInner().withPriority(110)
                    .withRuleCollections(Arrays
                        .asList(new FirewallPolicyFilterRuleCollection().withName("Example-Filter-Rule-Collection")
                            .withAction(new FirewallPolicyFilterRuleCollectionAction()
                                .withType(FirewallPolicyFilterRuleCollectionActionType.ALLOW))
                            .withRules(Arrays.asList(
                                new ApplicationRule().withName("rule1").withDescription("Insert trusted tenants header")
                                    .withSourceAddresses(Arrays.asList("216.58.216.164", "10.0.0.0/24"))
                                    .withProtocols(Arrays.asList(new FirewallPolicyRuleApplicationProtocol()
                                        .withProtocolType(FirewallPolicyRuleApplicationProtocolType.HTTP).withPort(80)))
                                    .withFqdnTags(Arrays.asList("WindowsVirtualDesktop"))
                                    .withHttpHeadersToInsert(Arrays.asList(new FirewallPolicyHttpHeaderToInsert()
                                        .withHeaderName("Restrict-Access-To-Tenants")
                                        .withHeaderValue("contoso.com,fabrikam.onmicrosoft.com"))))))),
                com.azure.core.util.Context.NONE);
    }
}
