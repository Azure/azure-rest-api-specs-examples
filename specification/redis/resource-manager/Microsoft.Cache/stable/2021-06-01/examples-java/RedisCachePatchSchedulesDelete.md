Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.redis.models.DefaultName;

/** Samples for PatchSchedules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2021-06-01/examples/RedisCachePatchSchedulesDelete.json
     */
    /**
     * Sample code: RedisCachePatchSchedulesDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCachePatchSchedulesDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .redisCaches()
            .manager()
            .serviceClient()
            .getPatchSchedules()
            .deleteWithResponse("rg1", "cache1", DefaultName.DEFAULT, Context.NONE);
    }
}
```
