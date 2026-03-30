
/**
 * Samples for StaticSites GetDatabaseConnections.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetStaticSiteDatabaseConnections.json
     */
    /**
     * Sample code: List overviews of database connections for the static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listOverviewsOfDatabaseConnectionsForTheStaticSite(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().getDatabaseConnections("rg", "testStaticSite0",
            com.azure.core.util.Context.NONE);
    }
}
