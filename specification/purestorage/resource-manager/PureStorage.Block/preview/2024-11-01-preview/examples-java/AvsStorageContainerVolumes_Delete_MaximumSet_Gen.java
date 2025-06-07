
/**
 * Samples for AvsStorageContainerVolumes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/AvsStorageContainerVolumes_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: AvsStorageContainerVolumes_Delete.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void
        avsStorageContainerVolumesDelete(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.avsStorageContainerVolumes().delete("rgpurestorage", "storagePoolname", "name", "cbdec-ddbb",
            com.azure.core.util.Context.NONE);
    }
}
