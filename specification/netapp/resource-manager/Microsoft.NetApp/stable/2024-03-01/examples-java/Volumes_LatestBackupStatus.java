
/**
 * Samples for Backups GetLatestStatus.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-03-01/examples/Volumes_LatestBackupStatus.json
     */
    /**
     * Sample code: Volumes_BackupStatus.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesBackupStatus(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.backups().getLatestStatusWithResponse("myRG", "account1", "pool1", "volume1",
            com.azure.core.util.Context.NONE);
    }
}
