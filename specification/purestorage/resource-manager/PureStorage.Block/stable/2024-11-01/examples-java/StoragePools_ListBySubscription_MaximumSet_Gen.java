
/**
 * Samples for StoragePools List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/StoragePools_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: StoragePools_ListBySubscription.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void
        storagePoolsListBySubscription(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.storagePools().list(com.azure.core.util.Context.NONE);
    }
}
