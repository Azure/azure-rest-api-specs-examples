
/**
 * Samples for Volumes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-07-01/examples/Volumes_Delete.json
     */
    /**
     * Sample code: Volumes_Delete.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumes().delete("myRG", "account1", "pool1", "volume1", null, com.azure.core.util.Context.NONE);
    }
}
