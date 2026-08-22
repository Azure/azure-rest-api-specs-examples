
/**
 * Samples for ServerEndpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/ServerEndpoints_Get.json
     */
    /**
     * Sample code: ServerEndpoints_Get.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void serverEndpointsGet(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.serverEndpoints().getWithResponse("SampleResourceGroup_1", "SampleStorageSyncService_1",
            "SampleSyncGroup_1", "SampleServerEndpoint_1", com.azure.core.util.Context.NONE);
    }
}
