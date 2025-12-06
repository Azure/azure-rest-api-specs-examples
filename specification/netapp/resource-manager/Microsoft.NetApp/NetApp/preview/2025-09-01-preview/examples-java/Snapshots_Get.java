
/**
 * Samples for Snapshots Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Snapshots_Get.json
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
