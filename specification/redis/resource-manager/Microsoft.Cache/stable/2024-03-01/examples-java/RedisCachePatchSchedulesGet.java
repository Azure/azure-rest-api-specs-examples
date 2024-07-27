
import com.azure.resourcemanager.redis.models.DefaultName;

/**
 * Samples for PatchSchedules Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redis/resource-manager/Microsoft.Cache/stable/2024-03-01/examples/RedisCachePatchSchedulesGet.json
     */
    /**
     * Sample code: RedisCachePatchSchedulesGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCachePatchSchedulesGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.redisCaches().manager().serviceClient().getPatchSchedules().getWithResponse("rg1", "cache1",
            DefaultName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
