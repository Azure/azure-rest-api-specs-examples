/** Samples for Backups GetStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-11-01/examples/Volumes_BackupStatus.json
     */
    /**
     * Sample code: Volumes_BackupStatus.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesBackupStatus(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager
            .backups()
            .getStatusWithResponse("myRG", "account1", "pool1", "volume1", com.azure.core.util.Context.NONE);
    }
}
