
/**
 * Samples for BlobContainers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/BlobContainersDelete.json
     */
    /**
     * Sample code: DeleteContainers.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void deleteContainers(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobContainers().deleteWithResponse("res4079", "sto4506", "container9689",
            com.azure.core.util.Context.NONE);
    }
}
