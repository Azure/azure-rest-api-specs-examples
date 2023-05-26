/** Samples for Redis ListUpgradeNotifications. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2023-04-01/examples/RedisCacheListUpgradeNotifications.json
     */
    /**
     * Sample code: RedisCacheListUpgradeNotifications.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheListUpgradeNotifications(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .redisCaches()
            .manager()
            .serviceClient()
            .getRedis()
            .listUpgradeNotifications("rg1", "cache1", 5000.0, com.azure.core.util.Context.NONE);
    }
}
