
/**
 * Samples for BlobServices GetServiceProperties.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/BlobServicesGet.json
     */
    /**
     * Sample code: GetBlobServices.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void getBlobServices(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobServices().getServicePropertiesWithResponse("res4410", "sto8607",
            com.azure.core.util.Context.NONE);
    }
}
