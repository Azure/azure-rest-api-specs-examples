```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.redisenterprise.models.ForceUnlinkParameters;
import java.util.Arrays;

/** Samples for Databases ForceUnlink. */
public final class Main {
    /*
     * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseDatabasesForceUnlink.json
     */
    /**
     * Sample code: How to unlink a database during a regional outage.
     *
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void howToUnlinkADatabaseDuringARegionalOutage(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager
            .databases()
            .forceUnlink(
                "rg1",
                "cache1",
                "default",
                new ForceUnlinkParameters()
                    .withIds(
                        Arrays
                            .asList(
                                "/subscriptions/subid2/resourceGroups/rg2/providers/Microsoft.Cache/redisEnterprise/cache2/databases/default")),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-redisenterprise_1.1.0-beta.1/sdk/redisenterprise/azure-resourcemanager-redisenterprise/README.md) on how to add the SDK to your project and authenticate.
