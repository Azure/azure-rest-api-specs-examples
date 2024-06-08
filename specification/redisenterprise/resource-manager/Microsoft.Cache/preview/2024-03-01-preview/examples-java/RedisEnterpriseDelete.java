
/**
 * Samples for RedisEnterprise Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2024-03-01-preview/examples/
     * RedisEnterpriseDelete.json
     */
    /**
     * Sample code: RedisEnterpriseDelete.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseDelete(com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.redisEnterprises().delete("rg1", "cache1", com.azure.core.util.Context.NONE);
    }
}
