```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.redis.models.ExportRdbParameters;

/** Samples for Redis ExportData. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2021-06-01/examples/RedisCacheExport.json
     */
    /**
     * Sample code: RedisCacheExport.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheExport(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .redisCaches()
            .manager()
            .serviceClient()
            .getRedis()
            .exportData(
                "rg1",
                "cache1",
                new ExportRdbParameters()
                    .withFormat("RDB")
                    .withPrefix("datadump1")
                    .withContainer("https://contosostorage.blob.core.window.net/urltoBlobContainer?sasKeyParameters"),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
