
/**
 * Samples for StaticSites GetDatabaseConnectionWithDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * GetStaticSiteDatabaseConnectionWithDetails.json
     */
    /**
     * Sample code: Get details of database connections for the static site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getDetailsOfDatabaseConnectionsForTheStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().getDatabaseConnectionWithDetailsWithResponse("rg",
            "testStaticSite0", "default", com.azure.core.util.Context.NONE);
    }
}
