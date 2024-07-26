
/**
 * Samples for AccessPolicy Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redis/resource-manager/Microsoft.Cache/stable/2024-03-01/examples/RedisCacheAccessPolicyDelete.json
     */
    /**
     * Sample code: RedisCacheAccessPolicyDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheAccessPolicyDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.redisCaches().manager().serviceClient().getAccessPolicies().delete("rg1", "cache1", "accessPolicy1",
            com.azure.core.util.Context.NONE);
    }
}
