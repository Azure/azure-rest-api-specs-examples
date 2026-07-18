
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListOperations.json
     */
    /**
     * Sample code: Lists all of the available SQL Rest API operations.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        listsAllOfTheAvailableSQLRestAPIOperations(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getOperations().list(com.azure.core.util.Context.NONE);
    }
}
