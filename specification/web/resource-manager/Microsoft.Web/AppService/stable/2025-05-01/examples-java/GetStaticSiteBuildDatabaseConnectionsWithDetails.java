
/**
 * Samples for StaticSites GetBuildDatabaseConnectionsWithDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetStaticSiteBuildDatabaseConnectionsWithDetails.json
     */
    /**
     * Sample code: List full details of database connections for the static site build.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listFullDetailsOfDatabaseConnectionsForTheStaticSiteBuild(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().getBuildDatabaseConnectionsWithDetails("rg", "testStaticSite0",
            "default", com.azure.core.util.Context.NONE);
    }
}
