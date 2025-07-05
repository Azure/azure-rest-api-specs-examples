
/**
 * Samples for AvsVmVolumes ListByAvsVm.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/AvsVmVolumes_ListByAvsVm_MaximumSet_Gen.json
     */
    /**
     * Sample code: AvsVmVolumes_ListByAvsVm.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void
        avsVmVolumesListByAvsVm(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.avsVmVolumes().listByAvsVm("rgpurestorage", "storagePoolname", "cbdec-ddbb",
            com.azure.core.util.Context.NONE);
    }
}
