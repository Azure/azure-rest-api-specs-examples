
import com.azure.resourcemanager.netapp.models.Volume;

/**
 * Samples for Volumes Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-07-01/examples/Volumes_Update.json
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
