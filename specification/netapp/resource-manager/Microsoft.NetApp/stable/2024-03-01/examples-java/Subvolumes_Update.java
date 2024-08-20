
import com.azure.resourcemanager.netapp.models.SubvolumeInfo;

/**
 * Samples for Subvolumes Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-03-01/examples/Subvolumes_Update.json
     */
    /**
     * Sample code: Subvolumes_Update.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void subvolumesUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        SubvolumeInfo resource = manager.subvolumes()
            .getWithResponse("myRG", "account1", "pool1", "volume1", "subvolume1", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withPath("/subvolumePath").apply();
    }
}
