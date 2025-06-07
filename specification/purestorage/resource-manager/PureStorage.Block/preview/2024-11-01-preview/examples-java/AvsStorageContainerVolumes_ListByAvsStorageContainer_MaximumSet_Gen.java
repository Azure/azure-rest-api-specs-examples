
/**
 * Samples for AvsStorageContainerVolumes ListByAvsStorageContainer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/AvsStorageContainerVolumes_ListByAvsStorageContainer_MaximumSet_Gen.json
     */
    /**
     * Sample code: AvsStorageContainerVolumes_ListByAvsStorageContainer.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void avsStorageContainerVolumesListByAvsStorageContainer(
        com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.avsStorageContainerVolumes().listByAvsStorageContainer("rgpurestorage", "storagePoolname", "name",
            com.azure.core.util.Context.NONE);
    }
}
