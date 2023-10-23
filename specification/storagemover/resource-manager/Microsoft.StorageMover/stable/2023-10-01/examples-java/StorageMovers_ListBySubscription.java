/** Samples for StorageMovers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2023-10-01/examples/StorageMovers_ListBySubscription.json
     */
    /**
     * Sample code: StorageMovers_List.
     *
     * @param manager Entry point to StorageMoverManager.
     */
    public static void storageMoversList(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.storageMovers().list(com.azure.core.util.Context.NONE);
    }
}
