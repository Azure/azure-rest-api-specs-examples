
/**
 * Samples for Subvolumes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-07-01/examples/Subvolumes_Delete.json
     */
    /**
     * Sample code: Subvolumes_Delete.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void subvolumesDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.subvolumes().delete("myRG", "account1", "pool1", "volume1", "subvolume1",
            com.azure.core.util.Context.NONE);
    }
}
