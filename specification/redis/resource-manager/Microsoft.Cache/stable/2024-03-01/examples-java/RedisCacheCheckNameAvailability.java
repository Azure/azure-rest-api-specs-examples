
import com.azure.resourcemanager.redis.models.CheckNameAvailabilityParameters;

/**
 * Samples for Redis CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redis/resource-manager/Microsoft.Cache/stable/2024-03-01/examples/RedisCacheCheckNameAvailability.
     * json
     */
    /**
     * Sample code: RedisCacheCheckNameAvailability.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheCheckNameAvailability(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.redisCaches().manager().serviceClient().getRedis().checkNameAvailabilityWithResponse(
            new CheckNameAvailabilityParameters().withName("cacheName").withType("Microsoft.Cache/Redis"),
            com.azure.core.util.Context.NONE);
    }
}
