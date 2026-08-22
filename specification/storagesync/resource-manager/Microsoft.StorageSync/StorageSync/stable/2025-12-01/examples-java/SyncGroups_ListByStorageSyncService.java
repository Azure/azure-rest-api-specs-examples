
/**
 * Samples for SyncGroups ListByStorageSyncService.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/SyncGroups_ListByStorageSyncService.json
     */
    /**
     * Sample code: SyncGroups_ListByStorageSyncService.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void
        syncGroupsListByStorageSyncService(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.syncGroups().listByStorageSyncService("SampleResourceGroup_1", "SampleStorageSyncService_1",
            com.azure.core.util.Context.NONE);
    }
}
