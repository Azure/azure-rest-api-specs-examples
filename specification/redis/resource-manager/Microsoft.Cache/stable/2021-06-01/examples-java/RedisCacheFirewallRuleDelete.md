Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for FirewallRules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2021-06-01/examples/RedisCacheFirewallRuleDelete.json
     */
    /**
     * Sample code: RedisCacheFirewallRuleDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheFirewallRuleDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .redisCaches()
            .manager()
            .serviceClient()
            .getFirewallRules()
            .deleteWithResponse("rg1", "cache1", "rule1", Context.NONE);
    }
}
```
