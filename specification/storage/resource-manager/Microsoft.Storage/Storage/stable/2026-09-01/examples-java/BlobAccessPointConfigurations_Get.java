
/**
 * Samples for BlobAccessPointConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/BlobAccessPointConfigurations_Get.json
     */
    /**
     * Sample code: GetBlobAccessPointConfiguration.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void getBlobAccessPointConfiguration(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobAccessPointConfigurations().getWithResponse("testrg", "teststorageaccount",
            "testaccesspointconfig", com.azure.core.util.Context.NONE);
    }
}
