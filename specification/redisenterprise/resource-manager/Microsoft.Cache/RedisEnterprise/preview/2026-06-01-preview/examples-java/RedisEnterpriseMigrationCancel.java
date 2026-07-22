
/**
 * Samples for Migrations Cancel.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/RedisEnterpriseMigrationCancel.json
     */
    /**
     * Sample code: RedisEnterpriseMigrationCancel.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void
        redisEnterpriseMigrationCancel(com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.migrations().cancel("rg1", "cache1", com.azure.core.util.Context.NONE);
    }
}
