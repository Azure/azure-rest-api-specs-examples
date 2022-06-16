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
