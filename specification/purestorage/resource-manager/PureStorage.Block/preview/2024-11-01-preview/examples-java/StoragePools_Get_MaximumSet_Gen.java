
/**
 * Samples for StoragePools GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/StoragePools_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: StoragePools_Get.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void storagePoolsGet(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.storagePools().getByResourceGroupWithResponse("rgpurestorage", "storagePoolname",
            com.azure.core.util.Context.NONE);
    }
}
