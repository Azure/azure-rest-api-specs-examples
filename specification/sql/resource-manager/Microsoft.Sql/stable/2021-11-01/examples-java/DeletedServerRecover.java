
/**
 * Samples for DeletedServers Recover.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DeletedServerRecover.json
     */
    /**
     * Sample code: Recover deleted server.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void recoverDeletedServer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDeletedServers().recover("japaneast", "sqlcrudtest-d-1414",
            com.azure.core.util.Context.NONE);
    }
}
