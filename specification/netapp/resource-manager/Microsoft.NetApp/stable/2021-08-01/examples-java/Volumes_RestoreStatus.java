import com.azure.core.util.Context;

/** Samples for Backups GetVolumeRestoreStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-08-01/examples/Volumes_RestoreStatus.json
     */
    /**
     * Sample code: Volumes_RestoreStatus.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesRestoreStatus(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.backups().getVolumeRestoreStatusWithResponse("myRG", "account1", "pool1", "volume1", Context.NONE);
    }
}
