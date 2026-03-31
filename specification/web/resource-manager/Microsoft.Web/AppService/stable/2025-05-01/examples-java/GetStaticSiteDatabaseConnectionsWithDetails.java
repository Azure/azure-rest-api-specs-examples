
/**
 * Samples for StaticSites GetDatabaseConnectionsWithDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetStaticSiteDatabaseConnectionsWithDetails.json
     */
    /**
     * Sample code: List full details of database connections for the static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listFullDetailsOfDatabaseConnectionsForTheStaticSite(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().getDatabaseConnectionsWithDetails("rg", "testStaticSite0",
            com.azure.core.util.Context.NONE);
    }
}
