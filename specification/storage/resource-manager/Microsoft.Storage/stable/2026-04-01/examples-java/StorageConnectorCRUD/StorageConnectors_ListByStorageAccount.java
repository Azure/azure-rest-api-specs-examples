
/**
 * Samples for Connectors ListByStorageAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/StorageConnectorCRUD/StorageConnectors_ListByStorageAccount.json
     */
    /**
     * Sample code: ListConnectorsByStorageAccount.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void listConnectorsByStorageAccount(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getConnectors().listByStorageAccount("testrg", "teststorageaccount",
            com.azure.core.util.Context.NONE);
    }
}
