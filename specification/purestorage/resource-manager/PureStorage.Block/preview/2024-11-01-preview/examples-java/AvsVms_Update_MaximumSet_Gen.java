
import com.azure.resourcemanager.purestorageblock.models.AvsVmUpdate;
import com.azure.resourcemanager.purestorageblock.models.AvsVmUpdateProperties;
import com.azure.resourcemanager.purestorageblock.models.SoftDeletion;

/**
 * Samples for AvsVms Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/AvsVms_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: AvsVms_Update.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void avsVmsUpdate(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.avsVms().update("rgpurestorage", "storagePoolname", "cbdec-ddbb",
            new AvsVmUpdate()
                .withProperties(new AvsVmUpdateProperties().withSoftDeletion(new SoftDeletion().withDestroyed(true))),
            com.azure.core.util.Context.NONE);
    }
}
