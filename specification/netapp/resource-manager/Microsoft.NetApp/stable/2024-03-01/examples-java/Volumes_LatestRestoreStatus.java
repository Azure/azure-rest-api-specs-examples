
/**
 * Samples for Backups GetVolumeLatestRestoreStatus.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-03-01/examples/Volumes_LatestRestoreStatus.
     * json
     */
    /**
     * Sample code: Volumes_RestoreStatus.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesRestoreStatus(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.backups().getVolumeLatestRestoreStatusWithResponse("myRG", "account1", "pool1", "volume1",
            com.azure.core.util.Context.NONE);
    }
}
