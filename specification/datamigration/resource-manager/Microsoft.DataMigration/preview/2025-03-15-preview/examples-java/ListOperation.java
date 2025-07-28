
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/
     * ListOperation.json
     */
    /**
     * Sample code: Lists all of the available SQL Rest API operations.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void listsAllOfTheAvailableSQLRestAPIOperations(
        com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
