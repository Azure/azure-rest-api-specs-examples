import com.azure.core.util.Context;

/** Samples for Redis GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2022-06-01/examples/RedisCacheGet.json
     */
    /**
     * Sample code: RedisCacheGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .redisCaches()
            .manager()
            .serviceClient()
            .getRedis()
            .getByResourceGroupWithResponse("rg1", "cache1", Context.NONE);
    }
}
