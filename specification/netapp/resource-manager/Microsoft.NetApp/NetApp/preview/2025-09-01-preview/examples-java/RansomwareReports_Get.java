
/**
 * Samples for RansomwareReports Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/RansomwareReports_Get.json
     */
    /**
     * Sample code: RansomwareReports_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void ransomwareReportsGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.ransomwareReports().getWithResponse("myRG", "account1", "pool1", "volume1", "ransomwareReport1",
            com.azure.core.util.Context.NONE);
    }
}
