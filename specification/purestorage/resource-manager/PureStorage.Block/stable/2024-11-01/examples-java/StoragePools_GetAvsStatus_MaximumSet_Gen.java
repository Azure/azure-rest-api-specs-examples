
/**
 * Samples for StoragePools GetAvsStatus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/StoragePools_GetAvsStatus_MaximumSet_Gen.json
     */
    /**
     * Sample code: StoragePools_GetAvsStatus.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void
        storagePoolsGetAvsStatus(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.storagePools().getAvsStatusWithResponse("rgpurestorage", "storagePoolname",
            com.azure.core.util.Context.NONE);
    }
}
