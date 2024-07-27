
import com.azure.resourcemanager.redis.models.ImportRdbParameters;
import java.util.Arrays;

/**
 * Samples for Redis ImportData.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redis/resource-manager/Microsoft.Cache/stable/2024-03-01/examples/RedisCacheImport.json
     */
    /**
     * Sample code: RedisCacheImport.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheImport(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.redisCaches().manager().serviceClient().getRedis().importData("rg1", "cache1",
            new ImportRdbParameters().withFormat("RDB")
                .withFiles(Arrays.asList("http://fileuris.contoso.com/pathtofile1"))
                .withStorageSubscriptionId("storageSubId"),
            com.azure.core.util.Context.NONE);
    }
}
