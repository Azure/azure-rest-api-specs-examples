
/**
 * Samples for AvsStorageContainerVolumes Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/AvsStorageContainerVolumes_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: AvsStorageContainerVolumes_Get.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void
        avsStorageContainerVolumesGet(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.avsStorageContainerVolumes().getWithResponse("rgpurestorage", "storagePoolname", "name", "cbdec-ddbb",
            com.azure.core.util.Context.NONE);
    }
}
