
/**
 * Samples for BlobServices List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/BlobServicesList.json
     */
    /**
     * Sample code: ListBlobServices.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void listBlobServices(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobServices().list("res4410", "sto8607", com.azure.core.util.Context.NONE);
    }
}
