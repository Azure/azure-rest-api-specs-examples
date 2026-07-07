
/**
 * Samples for BlobContainers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/BlobContainersGet.json
     */
    /**
     * Sample code: GetContainers.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void getContainers(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobContainers().getWithResponse("res9871", "sto6217", "container1634",
            com.azure.core.util.Context.NONE);
    }
}
