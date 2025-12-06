
/**
 * Samples for Subvolumes Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Subvolumes_Get.json
     */
    /**
     * Sample code: Subvolumes_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void subvolumesGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.subvolumes().getWithResponse("myRG", "account1", "pool1", "volume1", "subvolume1",
            com.azure.core.util.Context.NONE);
    }
}
