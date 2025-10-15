
/**
 * Samples for Buckets Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-07-01-preview/examples/Buckets_Get.json
     */
    /**
     * Sample code: Buckets_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void bucketsGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.buckets().getWithResponse("myRG", "account1", "pool1", "volume1", "bucket1",
            com.azure.core.util.Context.NONE);
    }
}
