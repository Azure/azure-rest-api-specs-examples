
/**
 * Samples for ServerOperations ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ListServerOperations.json
     */
    /**
     * Sample code: List the server management operations.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listTheServerManagementOperations(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerOperations().listByServer("sqlcrudtest-7398",
            "sqlcrudtest-4645", com.azure.core.util.Context.NONE);
    }
}
