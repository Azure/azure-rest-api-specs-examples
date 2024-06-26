import com.azure.resourcemanager.redisenterprise.models.FlushParameters;
import java.util.Arrays;

/** Samples for Databases Flush. */
public final class Main {
    /*
     * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2023-03-01-preview/examples/RedisEnterpriseDatabasesFlush.json
     */
    /**
     * Sample code: How to flush all the keys in the database.
     *
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void howToFlushAllTheKeysInTheDatabase(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager
            .databases()
            .flush(
                "rg1",
                "cache1",
                "default",
                new FlushParameters()
                    .withIds(
                        Arrays
                            .asList(
                                "/subscriptions/subid2/resourceGroups/rg2/providers/Microsoft.Cache/redisEnterprise/cache2/databases/default")),
                com.azure.core.util.Context.NONE);
    }
}
