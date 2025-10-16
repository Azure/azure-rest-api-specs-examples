
/**
 * Samples for Snapshots Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-07-01-preview/examples/Snapshots_Get.json
     */
    /**
     * Sample code: Snapshots_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void snapshotsGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.snapshots().getWithResponse("myRG", "account1", "pool1", "volume1", "snapshot1",
            com.azure.core.util.Context.NONE);
    }
}
