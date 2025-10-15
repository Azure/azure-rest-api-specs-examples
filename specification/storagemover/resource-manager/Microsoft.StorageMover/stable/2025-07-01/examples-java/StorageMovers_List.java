
/**
 * Samples for StorageMovers ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/StorageMovers_List.json
     */
    /**
     * Sample code: StorageMovers_List.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void storageMoversList(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.storageMovers().listByResourceGroup("examples-rg", com.azure.core.util.Context.NONE);
    }
}
