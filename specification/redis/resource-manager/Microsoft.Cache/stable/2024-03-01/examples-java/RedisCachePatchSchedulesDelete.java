
import com.azure.resourcemanager.redis.models.DefaultName;

/**
 * Samples for PatchSchedules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redis/resource-manager/Microsoft.Cache/stable/2024-03-01/examples/RedisCachePatchSchedulesDelete.
     * json
     */
    /**
     * Sample code: RedisCachePatchSchedulesDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCachePatchSchedulesDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.redisCaches().manager().serviceClient().getPatchSchedules().deleteWithResponse("rg1", "cache1",
            DefaultName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
