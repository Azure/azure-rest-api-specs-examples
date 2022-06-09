```java
import com.azure.core.util.Context;

/** Samples for FirewallPolicyRuleCollectionGroups List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/FirewallPolicyRuleCollectionGroupList.json
     */
    /**
     * Sample code: List all FirewallPolicyRuleCollectionGroups for a given FirewallPolicy.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllFirewallPolicyRuleCollectionGroupsForAGivenFirewallPolicy(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getFirewallPolicyRuleCollectionGroups()
            .list("rg1", "firewallPolicy", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
