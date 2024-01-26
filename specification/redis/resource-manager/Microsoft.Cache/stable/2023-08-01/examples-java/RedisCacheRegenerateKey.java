
import com.azure.resourcemanager.redis.models.RedisKeyType;
import com.azure.resourcemanager.redis.models.RedisRegenerateKeyParameters;

/** Samples for Redis RegenerateKey. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redis/resource-manager/Microsoft.Cache/stable/2023-08-01/examples/RedisCacheRegenerateKey.json
     */
    /**
     * Sample code: RedisCacheRegenerateKey.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheRegenerateKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.redisCaches().manager().serviceClient().getRedis().regenerateKeyWithResponse("rg1", "cache1",
            new RedisRegenerateKeyParameters().withKeyType(RedisKeyType.PRIMARY), com.azure.core.util.Context.NONE);
    }
}
