
/**
 * Samples for RedisEnterprise ListSkusForScaling.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/RedisEnterpriseListSkusForScaling.json
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
