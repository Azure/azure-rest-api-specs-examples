import com.azure.core.util.Context;

/** Samples for FirewallRules List. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2022-06-01/examples/RedisCacheFirewallRulesList.json
     */
    /**
     * Sample code: RedisCacheFirewallRulesList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheFirewallRulesList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.redisCaches().manager().serviceClient().getFirewallRules().list("rg1", "cache1", Context.NONE);
    }
}
