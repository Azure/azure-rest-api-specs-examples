
/** Samples for Redis List. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redis/resource-manager/Microsoft.Cache/stable/2023-08-01/examples/RedisCacheList.json
     */
    /**
     * Sample code: RedisCacheList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.redisCaches().manager().serviceClient().getRedis().list(com.azure.core.util.Context.NONE);
    }
}
