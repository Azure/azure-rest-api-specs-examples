import com.azure.core.util.Context;

/** Samples for FirewallRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2021-06-01/examples/RedisCacheFirewallRuleGet.json
     */
    /**
     * Sample code: RedisCacheFirewallRuleGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheFirewallRuleGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .redisCaches()
            .manager()
            .serviceClient()
            .getFirewallRules()
            .getWithResponse("rg1", "cache1", "rule1", Context.NONE);
    }
}
