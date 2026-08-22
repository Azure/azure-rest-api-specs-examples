
/**
 * Samples for SyncGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/SyncGroups_Delete.json
     */
    /**
     * Sample code: SyncGroups_Delete.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void syncGroupsDelete(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.syncGroups().deleteWithResponse("SampleResourceGroup_1", "SampleStorageSyncService_1",
            "SampleSyncGroup_1", com.azure.core.util.Context.NONE);
    }
}
