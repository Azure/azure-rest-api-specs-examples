
/**
 * Samples for StoragePools DisableAvsConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/StoragePools_DisableAvsConnection_MaximumSet_Gen.json
     */
    /**
     * Sample code: StoragePools_DisableAvsConnection.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void
        storagePoolsDisableAvsConnection(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.storagePools().disableAvsConnection("rgpurestorage", "storagePoolname",
            com.azure.core.util.Context.NONE);
    }
}
