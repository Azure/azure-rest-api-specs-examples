
/**
 * Samples for Migrations UnlinkMigratedEndpoint.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/RedisEnterpriseMigrationUnlinkMigratedEndpoint.json
     */
    /**
     * Sample code: RedisEnterpriseMigrationUnlinkMigratedEndpoint.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseMigrationUnlinkMigratedEndpoint(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.migrations().unlinkMigratedEndpoint("rg1", "cache1", com.azure.core.util.Context.NONE);
    }
}
