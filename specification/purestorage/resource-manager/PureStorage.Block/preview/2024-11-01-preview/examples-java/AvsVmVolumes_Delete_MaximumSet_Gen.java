
/**
 * Samples for AvsVmVolumes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/AvsVmVolumes_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: AvsVmVolumes_Delete.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void avsVmVolumesDelete(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.avsVmVolumes().delete("rgpurestorage", "storagePoolname", "cbdec-ddbb", "cbdec-ddbb",
            com.azure.core.util.Context.NONE);
    }
}
