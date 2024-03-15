
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2023-11-01/examples/
     * RedisEnterpriseDeletePrivateEndpointConnection.json
     */
    /**
     * Sample code: RedisEnterpriseDeletePrivateEndpointConnection.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseDeletePrivateEndpointConnection(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.privateEndpointConnections().delete("rg1", "cache1", "pectest01", com.azure.core.util.Context.NONE);
    }
}
