Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.redis.models.RedisLinkedServerCreateParameters;
import com.azure.resourcemanager.redis.models.ReplicationRole;

/** Samples for LinkedServer Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2021-06-01/examples/RedisCacheLinkedServer_Create.json
     */
    /**
     * Sample code: LinkedServer_Create.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void linkedServerCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .redisCaches()
            .manager()
            .serviceClient()
            .getLinkedServers()
            .create(
                "rg1",
                "cache1",
                "cache2",
                new RedisLinkedServerCreateParameters()
                    .withLinkedRedisCacheId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/Redis/cache2")
                    .withLinkedRedisCacheLocation("West US")
                    .withServerRole(ReplicationRole.SECONDARY),
                Context.NONE);
    }
}
```
