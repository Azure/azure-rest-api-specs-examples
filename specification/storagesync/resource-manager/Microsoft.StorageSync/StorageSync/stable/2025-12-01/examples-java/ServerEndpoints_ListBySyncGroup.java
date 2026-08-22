
/**
 * Samples for ServerEndpoints ListBySyncGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/ServerEndpoints_ListBySyncGroup.json
     */
    /**
     * Sample code: ServerEndpoints_ListBySyncGroup.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void
        serverEndpointsListBySyncGroup(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.serverEndpoints().listBySyncGroup("SampleResourceGroup_1", "SampleStorageSyncService_1",
            "SampleSyncGroup_1", com.azure.core.util.Context.NONE);
    }
}
