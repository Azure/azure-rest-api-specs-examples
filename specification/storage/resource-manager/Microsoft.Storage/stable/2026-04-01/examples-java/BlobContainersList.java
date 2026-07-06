
/**
 * Samples for BlobContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/BlobContainersList.json
     */
    /**
     * Sample code: ListContainers.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void listContainers(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobContainers().list("res9290", "sto1590", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
