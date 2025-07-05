
import com.azure.resourcemanager.purestorageblock.models.AvsStorageContainerVolumeUpdate;
import com.azure.resourcemanager.purestorageblock.models.AvsStorageContainerVolumeUpdateProperties;
import com.azure.resourcemanager.purestorageblock.models.SoftDeletion;

/**
 * Samples for AvsStorageContainerVolumes Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/AvsStorageContainerVolumes_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: AvsStorageContainerVolumes_Update.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void
        avsStorageContainerVolumesUpdate(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.avsStorageContainerVolumes().update("rgpurestorage", "storagePoolname", "name", "cbdec-ddbb",
            new AvsStorageContainerVolumeUpdate().withProperties(new AvsStorageContainerVolumeUpdateProperties()
                .withSoftDeletion(new SoftDeletion().withDestroyed(true))),
            com.azure.core.util.Context.NONE);
    }
}
