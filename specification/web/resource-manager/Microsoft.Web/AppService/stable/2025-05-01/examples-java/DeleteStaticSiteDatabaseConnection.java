
/**
 * Samples for StaticSites DeleteDatabaseConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/DeleteStaticSiteDatabaseConnection.json
     */
    /**
     * Sample code: Delete a database connection from a static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        deleteADatabaseConnectionFromAStaticSite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().deleteDatabaseConnectionWithResponse("rg", "testStaticSite0",
            "default", com.azure.core.util.Context.NONE);
    }
}
