
/**
 * Samples for Volumes List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-03-01/examples/Volumes_List.json
     */
    /**
     * Sample code: Volumes_List.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumes().list("myRG", "account1", "pool1", com.azure.core.util.Context.NONE);
    }
}
