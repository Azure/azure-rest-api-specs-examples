
/**
 * Samples for Connectors Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/StorageConnectorCRUD/StorageConnectors_Delete.json
     */
    /**
     * Sample code: DeleteConnector.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void deleteConnector(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getConnectors().delete("testrg", "teststorageaccount", "testconnector",
            com.azure.core.util.Context.NONE);
    }
}
