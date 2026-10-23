
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/PrivateEndpointConnections_Get.json
     */
    /**
     * Sample code: Get a private endpoint connection.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void getAPrivateEndpointConnection(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.privateEndpointConnections().getWithResponse("example-resource-group", "example-cluster",
            "example-private-endpoint-connection", com.azure.core.util.Context.NONE);
    }
}
