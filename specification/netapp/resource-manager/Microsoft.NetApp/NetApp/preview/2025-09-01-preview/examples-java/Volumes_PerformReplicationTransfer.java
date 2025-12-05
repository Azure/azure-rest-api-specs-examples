
/**
 * Samples for Volumes PerformReplicationTransfer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Volumes_PerformReplicationTransfer.json
     */
    /**
     * Sample code: Volumes_PerformReplicationTransfer.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesPerformReplicationTransfer(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumes().performReplicationTransfer("myRG", "account1", "pool1", "volume1",
            com.azure.core.util.Context.NONE);
    }
}
