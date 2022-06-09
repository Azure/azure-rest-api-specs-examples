```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.redis.fluent.models.RedisFirewallRuleInner;

/** Samples for FirewallRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2021-06-01/examples/RedisCacheFirewallRuleCreate.json
     */
    /**
     * Sample code: RedisCacheFirewallRuleCreate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheFirewallRuleCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .redisCaches()
            .manager()
            .serviceClient()
            .getFirewallRules()
            .createOrUpdateWithResponse(
                "rg1",
                "cache1",
                "rule1",
                new RedisFirewallRuleInner().withStartIp("192.168.1.1").withEndIp("192.168.1.4"),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
