/** Samples for Redis Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2023-08-01/examples/RedisCacheDelete.json
     */
    /**
     * Sample code: RedisCacheDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .redisCaches()
            .manager()
            .serviceClient()
            .getRedis()
            .delete("rg1", "cache1", com.azure.core.util.Context.NONE);
    }
}
