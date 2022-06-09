```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.redisenterprise.models.ExportClusterParameters;

/** Samples for Databases Export. */
public final class Main {
    /*
     * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseDatabasesExport.json
     */
    /**
     * Sample code: RedisEnterpriseDatabasesExport.
     *
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseDatabasesExport(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager
            .databases()
            .export(
                "rg1",
                "cache1",
                "default",
                new ExportClusterParameters()
                    .withSasUri("https://contosostorage.blob.core.window.net/urlToBlobContainer?sasKeyParameters"),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-redisenterprise_1.1.0-beta.1/sdk/redisenterprise/azure-resourcemanager-redisenterprise/README.md) on how to add the SDK to your project and authenticate.
