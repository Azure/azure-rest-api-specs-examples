
/** Samples for AccessPolicy Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redis/resource-manager/Microsoft.Cache/stable/2023-08-01/examples/RedisCacheAccessPolicyGet.json
     */
    /**
     * Sample code: RedisCacheAccessPolicyGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheAccessPolicyGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.redisCaches().manager().serviceClient().getAccessPolicies().getWithResponse("rg1", "cache1",
            "accessPolicy1", com.azure.core.util.Context.NONE);
    }
}
