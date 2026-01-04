
/**
 * Samples for StaticSites GetDatabaseConnections.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * GetStaticSiteDatabaseConnections.json
     */
    /**
     * Sample code: List overviews of database connections for the static site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listOverviewsOfDatabaseConnectionsForTheStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().getDatabaseConnections("rg", "testStaticSite0",
            com.azure.core.util.Context.NONE);
    }
}
