
/**
 * Samples for AvsVms ListByStoragePool.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/AvsVms_ListByStoragePool_MaximumSet_Gen.json
     */
    /**
     * Sample code: AvsVms_ListByStoragePool.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void
        avsVmsListByStoragePool(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.avsVms().listByStoragePool("rgpurestorage", "storagePoolname", com.azure.core.util.Context.NONE);
    }
}
