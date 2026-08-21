
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/PrivateEndpointConnections_Get.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Get.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void privateEndpointConnectionsGet(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.privateEndpointConnections().getWithResponse("res6977", "sss2527", "{privateEndpointConnectionName}",
            com.azure.core.util.Context.NONE);
    }
}
