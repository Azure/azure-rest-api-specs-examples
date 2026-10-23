
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/PrivateEndpointConnections_GetDisconnectedAfterLastPoolDeletion.json
     */
    /**
     * Sample code: Get a private endpoint connection disconnected after the last pool was deleted.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void getAPrivateEndpointConnectionDisconnectedAfterTheLastPoolWasDeleted(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.privateEndpointConnections().getWithResponse("example-resource-group", "example-cluster",
            "example-private-endpoint-connection", com.azure.core.util.Context.NONE);
    }
}
