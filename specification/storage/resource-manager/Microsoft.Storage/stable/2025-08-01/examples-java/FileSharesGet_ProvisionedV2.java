
/**
 * Samples for FileShares Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/FileSharesGet_ProvisionedV2.json
     */
    /**
     * Sample code: GetShareProvisionedV2.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void getShareProvisionedV2(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getFileShares().getWithResponse("res9871", "sto6217", "share1634", null, null,
            com.azure.core.util.Context.NONE);
    }
}
