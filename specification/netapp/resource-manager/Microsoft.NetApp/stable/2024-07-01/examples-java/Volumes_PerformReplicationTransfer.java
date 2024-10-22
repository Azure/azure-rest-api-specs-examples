
/**
 * Samples for Volumes PerformReplicationTransfer.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-07-01/examples/
     * Volumes_PerformReplicationTransfer.json
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
