import com.azure.core.util.Context;

/** Samples for Snapshots Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-01-01/examples/Snapshots_Get.json
     */
    /**
     * Sample code: Snapshots_Get.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void snapshotsGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.snapshots().getWithResponse("myRG", "account1", "pool1", "volume1", "snapshot1", Context.NONE);
    }
}
