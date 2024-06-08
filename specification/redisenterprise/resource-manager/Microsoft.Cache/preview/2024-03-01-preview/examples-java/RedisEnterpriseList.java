
/**
 * Samples for RedisEnterprise List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2024-03-01-preview/examples/
     * RedisEnterpriseList.json
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
