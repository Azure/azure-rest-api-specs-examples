
/**
 * Samples for AvsStorageContainers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/AvsStorageContainers_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: AvsStorageContainers_Delete.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void
        avsStorageContainersDelete(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.avsStorageContainers().delete("rgpurestorage", "storagePoolName", "storageContainerName",
            com.azure.core.util.Context.NONE);
    }
}
