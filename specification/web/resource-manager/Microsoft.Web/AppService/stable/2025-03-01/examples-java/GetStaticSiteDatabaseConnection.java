
/**
 * Samples for StaticSites GetDatabaseConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * GetStaticSiteDatabaseConnection.json
     */
    /**
     * Sample code: Get overview of database connections for the static site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getOverviewOfDatabaseConnectionsForTheStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().getDatabaseConnectionWithResponse("rg",
            "testStaticSite0", "default", com.azure.core.util.Context.NONE);
    }
}
