/** Samples for Snapshots Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-05-01/examples/Snapshots_Delete.json
     */
    /**
     * Sample code: Snapshots_Delete.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void snapshotsDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager
            .snapshots()
            .delete("myRG", "account1", "pool1", "volume1", "snapshot1", com.azure.core.util.Context.NONE);
    }
}
