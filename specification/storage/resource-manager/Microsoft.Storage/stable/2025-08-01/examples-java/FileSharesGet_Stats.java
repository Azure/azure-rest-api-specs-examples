
/**
 * Samples for FileShares Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/FileSharesGet_Stats.json
     */
    /**
     * Sample code: GetShareStats.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void getShareStats(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getFileShares().getWithResponse("res9871", "sto6217", "share1634", "stats", null,
            com.azure.core.util.Context.NONE);
    }
}
