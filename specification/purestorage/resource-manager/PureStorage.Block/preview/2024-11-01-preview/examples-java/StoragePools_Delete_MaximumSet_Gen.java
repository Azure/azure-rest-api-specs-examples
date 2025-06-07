
/**
 * Samples for StoragePools Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/StoragePools_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: StoragePools_Delete.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void storagePoolsDelete(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.storagePools().delete("rgpurestorage", "storagePoolname", com.azure.core.util.Context.NONE);
    }
}
