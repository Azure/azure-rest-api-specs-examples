
/**
 * Samples for CloudEndpoints RestoreHeartbeat.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/CloudEndpoints_RestoreHeatbeat.json
     */
    /**
     * Sample code: CloudEndpoints_restoreheartbeat.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void
        cloudEndpointsRestoreheartbeat(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.cloudEndpoints().restoreHeartbeatWithResponse("SampleResourceGroup_1", "SampleStorageSyncService_1",
            "SampleSyncGroup_1", "SampleCloudEndpoint_1", com.azure.core.util.Context.NONE);
    }
}
