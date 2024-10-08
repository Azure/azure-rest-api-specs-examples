
/**
 * Samples for PrivateLinkResources ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/PrivateLinkResourcesList.json
     */
    /**
     * Sample code: Gets private link resources for SQL.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsPrivateLinkResourcesForSQL(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getPrivateLinkResources().listByServer("Default", "test-svr",
            com.azure.core.util.Context.NONE);
    }
}
