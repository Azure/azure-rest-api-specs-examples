
/**
 * Samples for PrivateEndpointConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2023-11-01/examples/
     * RedisEnterpriseListPrivateEndpointConnections.json
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
