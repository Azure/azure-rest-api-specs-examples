
/**
 * Samples for Connectors Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/StorageConnectorCRUD/StorageConnectors_Get.json
     */
    /**
     * Sample code: GetConnector.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void getConnector(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getConnectors().getWithResponse("testrg", "teststorageaccount", "testconnector",
            com.azure.core.util.Context.NONE);
    }
}
