
/**
 * Samples for CloudEndpoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/CloudEndpoints_Delete.json
     */
    /**
     * Sample code: CloudEndpoints_Delete.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void cloudEndpointsDelete(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.cloudEndpoints().delete("SampleResourceGroup_1", "SampleStorageSyncService_1", "SampleSyncGroup_1",
            "SampleCloudEndpoint_1", com.azure.core.util.Context.NONE);
    }
}
