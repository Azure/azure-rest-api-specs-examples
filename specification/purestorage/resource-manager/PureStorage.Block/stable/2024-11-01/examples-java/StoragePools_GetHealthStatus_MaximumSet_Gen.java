
/**
 * Samples for StoragePools GetHealthStatus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/StoragePools_GetHealthStatus_MaximumSet_Gen.json
     */
    /**
     * Sample code: StoragePools_GetHealthStatus.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void
        storagePoolsGetHealthStatus(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.storagePools().getHealthStatusWithResponse("rgpurestorage", "storagePoolname",
            com.azure.core.util.Context.NONE);
    }
}
