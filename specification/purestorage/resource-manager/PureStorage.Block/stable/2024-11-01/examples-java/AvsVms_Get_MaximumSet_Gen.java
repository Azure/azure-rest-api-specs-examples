
/**
 * Samples for AvsVms Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/AvsVms_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: AvsVms_Get.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void avsVmsGet(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.avsVms().getWithResponse("rgpurestorage", "storagePoolname", "cbdec-ddbb",
            com.azure.core.util.Context.NONE);
    }
}
