
/**
 * Samples for StorageSyncServices List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/StorageSyncServices_ListBySubscription.json
     */
    /**
     * Sample code: StorageSyncServices_ListBySubscription.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void
        storageSyncServicesListBySubscription(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.storageSyncServices().list(com.azure.core.util.Context.NONE);
    }
}
