
import com.azure.resourcemanager.purestorageblock.models.AvsVmVolumeUpdate;
import com.azure.resourcemanager.purestorageblock.models.AvsVmVolumeUpdateProperties;
import com.azure.resourcemanager.purestorageblock.models.SoftDeletion;

/**
 * Samples for AvsVmVolumes Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/AvsVmVolumes_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: AvsVmVolumes_Update.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void avsVmVolumesUpdate(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.avsVmVolumes().update("rgpurestorage", "storagePoolname", "cbdec-ddbb", "cbdec-ddbb",
            new AvsVmVolumeUpdate().withProperties(
                new AvsVmVolumeUpdateProperties().withSoftDeletion(new SoftDeletion().withDestroyed(true))),
            com.azure.core.util.Context.NONE);
    }
}
