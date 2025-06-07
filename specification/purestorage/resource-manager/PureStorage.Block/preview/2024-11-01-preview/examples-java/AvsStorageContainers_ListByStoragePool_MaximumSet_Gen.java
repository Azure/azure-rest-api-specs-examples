
/**
 * Samples for AvsStorageContainers ListByStoragePool.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/AvsStorageContainers_ListByStoragePool_MaximumSet_Gen.json
     */
    /**
     * Sample code: AvsStorageContainers_ListByStoragePool.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void avsStorageContainersListByStoragePool(
        com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.avsStorageContainers().listByStoragePool("rgpurestorage", "spName", com.azure.core.util.Context.NONE);
    }
}
