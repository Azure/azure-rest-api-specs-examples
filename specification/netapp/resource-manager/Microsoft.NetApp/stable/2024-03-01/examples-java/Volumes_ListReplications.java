
/**
 * Samples for Volumes ListReplications.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-03-01/examples/Volumes_ListReplications.json
     */
    /**
     * Sample code: Volumes_ListReplications.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesListReplications(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumes().listReplications("myRG", "account1", "pool1", "volume1", com.azure.core.util.Context.NONE);
    }
}
