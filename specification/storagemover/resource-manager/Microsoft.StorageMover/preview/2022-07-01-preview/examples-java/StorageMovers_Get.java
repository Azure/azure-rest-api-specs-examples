/** Samples for StorageMovers GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2022-07-01-preview/examples/StorageMovers_Get.json
     */
    /**
     * Sample code: StorageMovers_Get.
     *
     * @param manager Entry point to StorageMoverManager.
     */
    public static void storageMoversGet(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager
            .storageMovers()
            .getByResourceGroupWithResponse(
                "examples-rg", "examples-storageMoverName", com.azure.core.util.Context.NONE);
    }
}
