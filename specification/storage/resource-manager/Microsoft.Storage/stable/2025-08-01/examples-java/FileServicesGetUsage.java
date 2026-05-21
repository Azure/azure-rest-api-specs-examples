
/**
 * Samples for FileServices GetServiceUsage.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/FileServicesGetUsage.json
     */
    /**
     * Sample code: GetFileServiceUsage.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void getFileServiceUsage(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getFileServices().getServiceUsageWithResponse("res4410", "sto8607",
            com.azure.core.util.Context.NONE);
    }
}
