
/**
 * Samples for AvsVmVolumes Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/AvsVmVolumes_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: AvsVmVolumes_Get.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void avsVmVolumesGet(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.avsVmVolumes().getWithResponse("rgpurestorage", "storagePoolname", "cbdec-ddbb", "cbdec-ddbb",
            com.azure.core.util.Context.NONE);
    }
}
