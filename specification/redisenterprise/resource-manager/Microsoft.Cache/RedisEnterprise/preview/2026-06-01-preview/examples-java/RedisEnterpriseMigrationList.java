
/**
 * Samples for Migrations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/RedisEnterpriseMigrationList.json
     */
    /**
     * Sample code: RedisEnterpriseMigrationList.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void
        redisEnterpriseMigrationList(com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.migrations().list("rg1", "cache1", com.azure.core.util.Context.NONE);
    }
}
