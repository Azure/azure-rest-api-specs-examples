
/**
 * Samples for CloudEndpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/CloudEndpoints_Get.json
     */
    /**
     * Sample code: CloudEndpoints_Get.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void cloudEndpointsGet(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.cloudEndpoints().getWithResponse("SampleResourceGroup_1", "SampleStorageSyncService_1",
            "SampleSyncGroup_1", "SampleCloudEndpoint_1", com.azure.core.util.Context.NONE);
    }
}
