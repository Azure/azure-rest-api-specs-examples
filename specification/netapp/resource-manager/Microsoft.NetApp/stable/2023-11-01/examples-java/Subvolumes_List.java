
/**
 * Samples for Subvolumes ListByVolume.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-11-01/examples/Subvolumes_List.json
     */
    /**
     * Sample code: Subvolumes_List.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void subvolumesList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.subvolumes().listByVolume("myRG", "account1", "pool1", "volume1", com.azure.core.util.Context.NONE);
    }
}
