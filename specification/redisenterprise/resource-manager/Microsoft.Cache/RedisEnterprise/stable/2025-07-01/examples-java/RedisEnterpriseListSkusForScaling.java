
/**
 * Samples for RedisEnterprise ListSkusForScaling.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redisenterprise/resource-manager/Microsoft.Cache/RedisEnterprise/stable/2025-07-01/examples/
     * RedisEnterpriseListSkusForScaling.json
     */
    /**
     * Sample code: RedisEnterpriseListSkusForScaling.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void
        redisEnterpriseListSkusForScaling(com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.redisEnterprises().listSkusForScalingWithResponse("rg1", "cache1", com.azure.core.util.Context.NONE);
    }
}
