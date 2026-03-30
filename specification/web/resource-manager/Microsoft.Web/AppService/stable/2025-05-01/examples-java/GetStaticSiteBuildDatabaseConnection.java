
/**
 * Samples for StaticSites GetBuildDatabaseConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetStaticSiteBuildDatabaseConnection.json
     */
    /**
     * Sample code: Get overview of database connections for the static site build.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getOverviewOfDatabaseConnectionsForTheStaticSiteBuild(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().getBuildDatabaseConnectionWithResponse("rg", "testStaticSite0",
            "default", "default", com.azure.core.util.Context.NONE);
    }
}
