
/**
 * Samples for AvsStorageContainers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/AvsStorageContainers_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: AvsStorageContainers_Get.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void
        avsStorageContainersGet(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.avsStorageContainers().getWithResponse("rgpurestorage", "storagePoolName", "storageContainerName",
            com.azure.core.util.Context.NONE);
    }
}
