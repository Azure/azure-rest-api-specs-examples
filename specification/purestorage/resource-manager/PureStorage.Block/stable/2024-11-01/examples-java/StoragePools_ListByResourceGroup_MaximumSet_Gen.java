
/**
 * Samples for StoragePools ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/StoragePools_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: StoragePools_ListByResourceGroup.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void
        storagePoolsListByResourceGroup(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.storagePools().listByResourceGroup("rgpurestorage", com.azure.core.util.Context.NONE);
    }
}
