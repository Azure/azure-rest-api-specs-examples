
/**
 * Samples for FileServices GetServiceProperties.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/FileServicesGet.json
     */
    /**
     * Sample code: GetFileServices.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void getFileServices(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getFileServices().getServicePropertiesWithResponse("res4410", "sto8607",
            com.azure.core.util.Context.NONE);
    }
}
