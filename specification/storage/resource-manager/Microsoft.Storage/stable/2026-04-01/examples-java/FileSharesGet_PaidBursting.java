
/**
 * Samples for FileShares Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/FileSharesGet_PaidBursting.json
     */
    /**
     * Sample code: GetSharePaidBursting.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void getSharePaidBursting(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getFileShares().getWithResponse("res9871", "sto6217", "share1634", null, null,
            com.azure.core.util.Context.NONE);
    }
}
