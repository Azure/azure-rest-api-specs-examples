
import com.azure.resourcemanager.netapp.models.Volume;

/**
 * Samples for Volumes Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Volumes_Update.json
     */
    /**
     * Sample code: Volumes_Update.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        Volume resource = manager.volumes()
            .getWithResponse("myRG", "account1", "pool1", "volume1", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
    }
}
