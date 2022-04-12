Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Redis GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2021-06-01/examples/RedisCacheGet.json
     */
    /**
     * Sample code: RedisCacheGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .redisCaches()
            .manager()
            .serviceClient()
            .getRedis()
            .getByResourceGroupWithResponse("rg1", "cache1", Context.NONE);
    }
}
```
