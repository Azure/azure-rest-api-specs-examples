
/**
 * Samples for TableServices GetServiceProperties.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/TableServicesGet.json
     */
    /**
     * Sample code: TableServicesGet.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void tableServicesGet(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getTableServices().getServicePropertiesWithResponse("res4410", "sto8607",
            com.azure.core.util.Context.NONE);
    }
}
