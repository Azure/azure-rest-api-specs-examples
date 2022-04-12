Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.redis.models.CheckNameAvailabilityParameters;

/** Samples for Redis CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2021-06-01/examples/RedisCacheCheckNameAvailability.json
     */
    /**
     * Sample code: RedisCacheCheckNameAvailability.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheCheckNameAvailability(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .redisCaches()
            .manager()
            .serviceClient()
            .getRedis()
            .checkNameAvailabilityWithResponse(
                new CheckNameAvailabilityParameters().withName("cacheName").withType("Microsoft.Cache/Redis"),
                Context.NONE);
    }
}
```
