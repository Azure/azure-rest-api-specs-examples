
/**
 * Samples for Databases UpgradeDBRedisVersion.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redisenterprise/resource-manager/Microsoft.Cache/RedisEnterprise/stable/2025-07-01/examples/
     * RedisEnterpriseDatabasesUpgradeDBRedisVersion.json
     */
    /**
     * Sample code: How to upgrade your database Redis version.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void
        howToUpgradeYourDatabaseRedisVersion(com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.databases().upgradeDBRedisVersion("rg1", "cache1", "default", com.azure.core.util.Context.NONE);
    }
}
