import com.azure.core.util.Context;

/** Samples for Snapshots List. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-01-01/examples/Snapshots_List.json
     */
    /**
     * Sample code: Snapshots_List.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void snapshotsList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.snapshots().list("myRG", "account1", "pool1", "volume1", Context.NONE);
    }
}
