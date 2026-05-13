
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/PrivateEndpointConnections_Get.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Get.
     * 
     * @param manager Entry point to FileSharesManager.
     */
    public static void privateEndpointConnectionsGet(com.azure.resourcemanager.fileshares.FileSharesManager manager) {
        manager.privateEndpointConnections().getWithResponse("rgfileshares", "fileshare", "privateEndpointConnection1",
            com.azure.core.util.Context.NONE);
    }
}
