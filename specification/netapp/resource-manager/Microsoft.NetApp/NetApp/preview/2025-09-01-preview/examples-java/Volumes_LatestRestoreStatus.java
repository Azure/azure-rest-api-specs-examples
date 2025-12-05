
/**
 * Samples for Backups GetVolumeLatestRestoreStatus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Volumes_LatestRestoreStatus.json
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
