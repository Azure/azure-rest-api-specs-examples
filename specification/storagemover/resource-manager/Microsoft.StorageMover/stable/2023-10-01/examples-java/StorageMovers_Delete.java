/** Samples for StorageMovers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2023-10-01/examples/StorageMovers_Delete.json
     */
    /**
     * Sample code: StorageMovers_Delete.
     *
     * @param manager Entry point to StorageMoverManager.
     */
    public static void storageMoversDelete(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.storageMovers().delete("examples-rg", "examples-storageMoverName", com.azure.core.util.Context.NONE);
    }
}
