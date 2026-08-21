
/**
 * Samples for StorageSyncServices GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/StorageSyncServices_Get.json
     */
    /**
     * Sample code: StorageSyncServices_Get.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void storageSyncServicesGet(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.storageSyncServices().getByResourceGroupWithResponse("SampleResourceGroup_1",
            "SampleStorageSyncService_1", com.azure.core.util.Context.NONE);
    }
}
