import com.azure.core.util.Context;

/** Samples for Redis ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2022-06-01/examples/RedisCacheListKeys.json
     */
    /**
     * Sample code: RedisCacheListKeys.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheListKeys(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.redisCaches().manager().serviceClient().getRedis().listKeysWithResponse("rg1", "cache1", Context.NONE);
    }
}
