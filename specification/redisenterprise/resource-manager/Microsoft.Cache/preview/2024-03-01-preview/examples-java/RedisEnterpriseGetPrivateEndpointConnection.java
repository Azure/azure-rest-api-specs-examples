
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2024-03-01-preview/examples/
     * RedisEnterpriseGetPrivateEndpointConnection.json
     */
    /**
     * Sample code: RedisEnterpriseGetPrivateEndpointConnection.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseGetPrivateEndpointConnection(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.privateEndpointConnections().getWithResponse("rg1", "cache1", "pectest01",
            com.azure.core.util.Context.NONE);
    }
}
