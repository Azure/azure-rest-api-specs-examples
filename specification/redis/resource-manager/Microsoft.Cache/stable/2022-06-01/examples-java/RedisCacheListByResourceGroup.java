import com.azure.core.util.Context;

/** Samples for Redis ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2022-06-01/examples/RedisCacheListByResourceGroup.json
     */
    /**
     * Sample code: RedisCacheListByResourceGroup.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheListByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.redisCaches().manager().serviceClient().getRedis().listByResourceGroup("rg1", Context.NONE);
    }
}
