
/**
 * Samples for FirewallRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redis/resource-manager/Microsoft.Cache/stable/2024-03-01/examples/RedisCacheFirewallRuleDelete.json
     */
    /**
     * Sample code: RedisCacheFirewallRuleDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheFirewallRuleDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.redisCaches().manager().serviceClient().getFirewallRules().deleteWithResponse("rg1", "cache1", "rule1",
            com.azure.core.util.Context.NONE);
    }
}
