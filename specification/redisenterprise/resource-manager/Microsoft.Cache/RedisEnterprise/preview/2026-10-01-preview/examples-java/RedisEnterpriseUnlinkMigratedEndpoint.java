
/**
 * Samples for Clusters UnlinkMigratedEndpoint.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/RedisEnterpriseUnlinkMigratedEndpoint.json
     */
    /**
     * Sample code: RedisEnterpriseUnlinkMigratedEndpoint.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseUnlinkMigratedEndpoint(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.clusters().unlinkMigratedEndpoint("rg1", "cache1", com.azure.core.util.Context.NONE);
    }
}
