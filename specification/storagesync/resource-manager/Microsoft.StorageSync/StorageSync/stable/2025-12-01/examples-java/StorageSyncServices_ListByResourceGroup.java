
/**
 * Samples for StorageSyncServices ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/StorageSyncServices_ListByResourceGroup.json
     */
    /**
     * Sample code: StorageSyncServices_ListByResourceGroup.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void
        storageSyncServicesListByResourceGroup(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.storageSyncServices().listByResourceGroup("SampleResourceGroup_1", com.azure.core.util.Context.NONE);
    }
}
