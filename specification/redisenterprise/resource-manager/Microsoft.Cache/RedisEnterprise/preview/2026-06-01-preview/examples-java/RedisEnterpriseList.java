
/**
 * Samples for RedisEnterprise List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/RedisEnterpriseList.json
     */
    /**
     * Sample code: RedisEnterpriseList.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseList(com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.redisEnterprises().list(com.azure.core.util.Context.NONE);
    }
}
