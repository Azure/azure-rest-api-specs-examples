Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for FirewallPolicyRuleCollectionGroups Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/FirewallPolicyRuleCollectionGroupWithWebCategoriesGet.json
     */
    /**
     * Sample code: Get FirewallPolicyRuleCollectionGroup With Web Categories.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getFirewallPolicyRuleCollectionGroupWithWebCategories(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getFirewallPolicyRuleCollectionGroups()
            .getWithResponse("rg1", "firewallPolicy", "ruleCollectionGroup1", Context.NONE);
    }
}
```
