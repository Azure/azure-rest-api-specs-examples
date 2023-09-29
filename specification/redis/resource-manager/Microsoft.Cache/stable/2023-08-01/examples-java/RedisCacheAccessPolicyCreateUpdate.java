import com.azure.resourcemanager.redis.fluent.models.RedisCacheAccessPolicyInner;

/** Samples for AccessPolicy CreateUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2023-08-01/examples/RedisCacheAccessPolicyCreateUpdate.json
     */
    /**
     * Sample code: RedisCacheAccessPolicyCreateUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheAccessPolicyCreateUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .redisCaches()
            .manager()
            .serviceClient()
            .getAccessPolicies()
            .createUpdate(
                "rg1",
                "cache1",
                "accessPolicy1",
                new RedisCacheAccessPolicyInner().withPermissions("+get +hget"),
                com.azure.core.util.Context.NONE);
    }
}
