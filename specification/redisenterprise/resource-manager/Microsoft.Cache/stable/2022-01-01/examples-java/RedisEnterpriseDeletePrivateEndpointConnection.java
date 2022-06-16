import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseDeletePrivateEndpointConnection.json
     */
    /**
     * Sample code: RedisEnterpriseDeletePrivateEndpointConnection.
     *
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseDeletePrivateEndpointConnection(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.privateEndpointConnections().deleteWithResponse("rg1", "cache1", "pectest01", Context.NONE);
    }
}
