
/**
 * Samples for StoragePools RepairAvsConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/StoragePools_RepairAvsConnection_MaximumSet_Gen.json
     */
    /**
     * Sample code: StoragePools_RepairAvsConnection.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void
        storagePoolsRepairAvsConnection(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.storagePools().repairAvsConnection("rgpurestorage", "storagePoolname",
            com.azure.core.util.Context.NONE);
    }
}
