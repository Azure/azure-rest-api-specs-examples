
/**
 * Samples for StoragePools GetAvsConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/StoragePools_GetAvsConnection_MaximumSet_Gen.json
     */
    /**
     * Sample code: StoragePools_GetAvsConnection.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void
        storagePoolsGetAvsConnection(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.storagePools().getAvsConnectionWithResponse("rgpurestorage", "storagePoolname",
            com.azure.core.util.Context.NONE);
    }
}
