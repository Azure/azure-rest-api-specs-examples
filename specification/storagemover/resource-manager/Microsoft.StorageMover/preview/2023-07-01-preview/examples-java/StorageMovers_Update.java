import com.azure.resourcemanager.storagemover.models.StorageMover;

/** Samples for StorageMovers Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2023-07-01-preview/examples/StorageMovers_Update.json
     */
    /**
     * Sample code: StorageMovers_Update.
     *
     * @param manager Entry point to StorageMoverManager.
     */
    public static void storageMoversUpdate(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        StorageMover resource =
            manager
                .storageMovers()
                .getByResourceGroupWithResponse(
                    "examples-rg", "examples-storageMoverName", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withDescription("Updated Storage Mover Description").apply();
    }
}
