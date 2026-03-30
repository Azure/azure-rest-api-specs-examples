
/**
 * Samples for StaticSites DeleteBuildDatabaseConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/DeleteStaticSiteBuildDatabaseConnection.json
     */
    /**
     * Sample code: Delete a database connection from a static site build.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        deleteADatabaseConnectionFromAStaticSiteBuild(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().deleteBuildDatabaseConnectionWithResponse("rg", "testStaticSite0",
            "default", "default", com.azure.core.util.Context.NONE);
    }
}
