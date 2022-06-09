```java
import com.azure.core.util.Context;

/** Samples for Redis ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2021-06-01/examples/RedisCacheListKeys.json
     */
    /**
     * Sample code: RedisCacheListKeys.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheListKeys(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.redisCaches().manager().serviceClient().getRedis().listKeysWithResponse("rg1", "cache1", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
