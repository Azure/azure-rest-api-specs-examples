```java
import com.azure.core.util.Context;

/** Samples for LinkedServer Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2021-06-01/examples/RedisCacheLinkedServer_Get.json
     */
    /**
     * Sample code: LinkedServer_Get.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void linkedServerGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .redisCaches()
            .manager()
            .serviceClient()
            .getLinkedServers()
            .getWithResponse("rg1", "cache1", "cache2", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
