
/**
 * Samples for PrivateEndpointConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/RedisEnterpriseListPrivateEndpointConnections.json
     */
    /**
     * Sample code: RedisEnterpriseListPrivateEndpointConnections.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseListPrivateEndpointConnections(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.privateEndpointConnections().list("rg1", "cache1", com.azure.core.util.Context.NONE);
    }
}
