
/**
 * Samples for AvsVms Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/AvsVms_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: AvsVms_Delete.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void avsVmsDelete(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.avsVms().delete("rgpurestorage", "storagePoolname", "cbdec-ddbb", com.azure.core.util.Context.NONE);
    }
}
