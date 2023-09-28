import com.azure.resourcemanager.redis.models.ExportRdbParameters;

/** Samples for Redis ExportData. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2023-08-01/examples/RedisCacheExport.json
     */
    /**
     * Sample code: RedisCacheExport.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheExport(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .redisCaches()
            .manager()
            .serviceClient()
            .getRedis()
            .exportData(
                "rg1",
                "cache1",
                new ExportRdbParameters()
                    .withFormat("RDB")
                    .withPrefix("datadump1")
                    .withContainer("https://contosostorage.blob.core.window.net/urltoBlobContainer?sasKeyParameters")
                    .withStorageSubscriptionId("storageSubId"),
                com.azure.core.util.Context.NONE);
    }
}
