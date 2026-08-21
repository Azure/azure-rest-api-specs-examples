
/**
 * Samples for RegisteredServers ListByStorageSyncService.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/RegisteredServers_ListByStorageSyncService.json
     */
    /**
     * Sample code: RegisteredServers_ListByStorageSyncService.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void
        registeredServersListByStorageSyncService(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.registeredServers().listByStorageSyncService("SampleResourceGroup_1", "SampleStorageSyncService_1",
            com.azure.core.util.Context.NONE);
    }
}
