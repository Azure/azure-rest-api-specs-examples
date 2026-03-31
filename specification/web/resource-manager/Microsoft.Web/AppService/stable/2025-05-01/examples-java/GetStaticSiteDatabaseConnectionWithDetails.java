
/**
 * Samples for StaticSites GetDatabaseConnectionWithDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetStaticSiteDatabaseConnectionWithDetails.json
     */
    /**
     * Sample code: Get details of database connections for the static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getDetailsOfDatabaseConnectionsForTheStaticSite(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().getDatabaseConnectionWithDetailsWithResponse("rg", "testStaticSite0",
            "default", com.azure.core.util.Context.NONE);
    }
}
