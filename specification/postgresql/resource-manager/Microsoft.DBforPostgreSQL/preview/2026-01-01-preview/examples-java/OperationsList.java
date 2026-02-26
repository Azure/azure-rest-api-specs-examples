
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/OperationsList.json
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
