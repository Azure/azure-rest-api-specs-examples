
/**
 * Samples for ServerOperations ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListServerOperations.json
     */
    /**
     * Sample code: List the server management operations.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listTheServerManagementOperations(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerOperations().listByServer("sqlcrudtest-7398", "sqlcrudtest-4645",
            com.azure.core.util.Context.NONE);
    }
}
