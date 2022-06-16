import com.azure.core.util.Context;

/** Samples for Databases ListByCluster. */
public final class Main {
    /*
     * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseDatabasesListByCluster.json
     */
    /**
     * Sample code: RedisEnterpriseDatabasesListByCluster.
     *
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseDatabasesListByCluster(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.databases().listByCluster("rg1", "cache1", Context.NONE);
    }
}
