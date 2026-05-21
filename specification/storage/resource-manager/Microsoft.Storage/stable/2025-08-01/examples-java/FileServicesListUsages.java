
/**
 * Samples for FileServices ListServiceUsages.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/FileServicesListUsages.json
     */
    /**
     * Sample code: ListFileServiceUsages.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void listFileServiceUsages(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getFileServices().listServiceUsages("res4410", "sto8607", null,
            com.azure.core.util.Context.NONE);
    }
}
