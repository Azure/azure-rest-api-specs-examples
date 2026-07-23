
/**
 * Samples for Databases ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/RedisEnterpriseDatabasesListKeys.json
     */
    /**
     * Sample code: RedisEnterpriseDatabasesListKeys.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void
        redisEnterpriseDatabasesListKeys(com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.databases().listKeysWithResponse("rg1", "cache1", "default", com.azure.core.util.Context.NONE);
    }
}
