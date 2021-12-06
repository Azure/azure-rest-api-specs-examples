Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.FirewallPolicyRuleCollectionGroupInner;
import com.azure.resourcemanager.network.models.ApplicationRule;
import com.azure.resourcemanager.network.models.FirewallPolicyFilterRuleCollection;
import com.azure.resourcemanager.network.models.FirewallPolicyFilterRuleCollectionAction;
import com.azure.resourcemanager.network.models.FirewallPolicyFilterRuleCollectionActionType;
import com.azure.resourcemanager.network.models.FirewallPolicyRuleApplicationProtocol;
import com.azure.resourcemanager.network.models.FirewallPolicyRuleApplicationProtocolType;
import java.util.Arrays;

/** Samples for FirewallPolicyRuleCollectionGroups CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/FirewallPolicyRuleCollectionGroupWithWebCategoriesPut.json
     */
    /**
     * Sample code: Create FirewallPolicyRuleCollectionGroup With Web Categories.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createFirewallPolicyRuleCollectionGroupWithWebCategories(
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
                                                new ApplicationRule()
                                                    .withName("rule1")
                                                    .withDescription("Deny inbound rule")
                                                    .withSourceAddresses(Arrays.asList("216.58.216.164", "10.0.0.0/24"))
                                                    .withProtocols(
                                                        Arrays
                                                            .asList(
                                                                new FirewallPolicyRuleApplicationProtocol()
                                                                    .withProtocolType(
                                                                        FirewallPolicyRuleApplicationProtocolType.HTTPS)
                                                                    .withPort(443)))
                                                    .withWebCategories(Arrays.asList("Hacking")))))),
                Context.NONE);
    }
}
```
