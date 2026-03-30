
/**
 * Samples for StaticSites GetDatabaseConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetStaticSiteDatabaseConnection.json
     */
    /**
     * Sample code: Get overview of database connections for the static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getOverviewOfDatabaseConnectionsForTheStaticSite(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().getDatabaseConnectionWithResponse("rg", "testStaticSite0", "default",
            com.azure.core.util.Context.NONE);
    }
}
