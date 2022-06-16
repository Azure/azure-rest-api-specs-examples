import com.azure.core.util.Context;
import com.azure.resourcemanager.redis.models.RebootType;
import com.azure.resourcemanager.redis.models.RedisRebootParameters;
import java.util.Arrays;

/** Samples for Redis ForceReboot. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2021-06-01/examples/RedisCacheForceReboot.json
     */
    /**
     * Sample code: RedisCacheForceReboot.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheForceReboot(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .redisCaches()
            .manager()
            .serviceClient()
            .getRedis()
            .forceRebootWithResponse(
                "rg1",
                "cache1",
                new RedisRebootParameters()
                    .withRebootType(RebootType.ALL_NODES)
                    .withShardId(0)
                    .withPorts(Arrays.asList(13000, 15001)),
                Context.NONE);
    }
}
