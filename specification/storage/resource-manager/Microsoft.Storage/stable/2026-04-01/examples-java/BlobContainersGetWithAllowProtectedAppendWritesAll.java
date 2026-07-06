
/**
 * Samples for BlobContainers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/BlobContainersGetWithAllowProtectedAppendWritesAll.json
     */
    /**
     * Sample code: GetBlobContainersGetWithAllowProtectedAppendWritesAll.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void getBlobContainersGetWithAllowProtectedAppendWritesAll(
        com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobContainers().getWithResponse("res9871", "sto6217", "container1634",
            com.azure.core.util.Context.NONE);
    }
}
