```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.redis.models.ImportRdbParameters;
import java.util.Arrays;

/** Samples for Redis ImportData. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2021-06-01/examples/RedisCacheImport.json
     */
    /**
     * Sample code: RedisCacheImport.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheImport(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .redisCaches()
            .manager()
            .serviceClient()
            .getRedis()
            .importData(
                "rg1",
                "cache1",
                new ImportRdbParameters()
                    .withFormat("RDB")
                    .withFiles(Arrays.asList("http://fileuris.contoso.com/pathtofile1")),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
