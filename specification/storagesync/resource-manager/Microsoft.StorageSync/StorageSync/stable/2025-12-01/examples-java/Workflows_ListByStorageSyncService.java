
/**
 * Samples for Workflows ListByStorageSyncService.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/Workflows_ListByStorageSyncService.json
     */
    /**
     * Sample code: Workflows_ListByStorageSyncService.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void
        workflowsListByStorageSyncService(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.workflows().listByStorageSyncService("SampleResourceGroup_1", "SampleStorageSyncService_1",
            com.azure.core.util.Context.NONE);
    }
}
