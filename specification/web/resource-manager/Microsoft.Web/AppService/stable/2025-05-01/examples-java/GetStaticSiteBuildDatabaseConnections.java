
/**
 * Samples for StaticSites GetBuildDatabaseConnections.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetStaticSiteBuildDatabaseConnections.json
     */
    /**
     * Sample code: List overviews of database connections for the static site build.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listOverviewsOfDatabaseConnectionsForTheStaticSiteBuild(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().getBuildDatabaseConnections("rg", "testStaticSite0", "default",
            com.azure.core.util.Context.NONE);
    }
}
