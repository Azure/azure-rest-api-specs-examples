
/**
 * Samples for Redis FlushCache.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redis/resource-manager/Microsoft.Cache/stable/2024-03-01/examples/RedisCacheFlush.json
     */
    /**
     * Sample code: RedisCacheFlush.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheFlush(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.redisCaches().manager().serviceClient().getRedis().flushCache("resource-group-name", "cache-name",
            com.azure.core.util.Context.NONE);
    }
}
