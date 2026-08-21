
/**
 * Samples for StorageSyncServices Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/StorageSyncServices_Delete.json
     */
    /**
     * Sample code: StorageSyncServices_Delete.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void storageSyncServicesDelete(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.storageSyncServices().delete("SampleResourceGroup_1", "SampleStorageSyncService_1",
            com.azure.core.util.Context.NONE);
    }
}
