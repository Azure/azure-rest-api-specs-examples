
/**
 * Samples for Databases Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redisenterprise/resource-manager/Microsoft.Cache/RedisEnterprise/stable/2025-07-01/examples/
     * RedisEnterpriseDatabasesDelete.json
     */
    /**
     * Sample code: RedisEnterpriseDatabasesDelete.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void
        redisEnterpriseDatabasesDelete(com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.databases().delete("rg1", "cache1", "db1", com.azure.core.util.Context.NONE);
    }
}
