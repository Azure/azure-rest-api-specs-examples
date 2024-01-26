
import com.azure.resourcemanager.redis.fluent.models.RedisFirewallRuleInner;

/** Samples for FirewallRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redis/resource-manager/Microsoft.Cache/stable/2023-08-01/examples/RedisCacheFirewallRuleCreate.json
     */
    /**
     * Sample code: RedisCacheFirewallRuleCreate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheFirewallRuleCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.redisCaches().manager().serviceClient().getFirewallRules().createOrUpdateWithResponse("rg1", "cache1",
            "rule1", new RedisFirewallRuleInner().withStartIp("192.168.1.1").withEndIp("192.168.1.4"),
            com.azure.core.util.Context.NONE);
    }
}
