
/** Samples for PrivateLinkResources ListByRedisCache. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2023-08-01/examples/
     * RedisCacheListPrivateLinkResources.json
     */
    /**
     * Sample code: StorageAccountListPrivateLinkResources.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountListPrivateLinkResources(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.redisCaches().manager().serviceClient().getPrivateLinkResources().listByRedisCache("rgtest01",
            "cacheTest01", com.azure.core.util.Context.NONE);
    }
}
