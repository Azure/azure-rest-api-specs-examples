
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/OperationsList.
     * json
     */
    /**
     * Sample code: List all available REST API operations.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listAllAvailableRESTAPIOperations(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
