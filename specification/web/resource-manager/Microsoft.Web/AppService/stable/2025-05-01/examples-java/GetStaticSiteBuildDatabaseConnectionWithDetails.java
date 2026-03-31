
/**
 * Samples for StaticSites GetBuildDatabaseConnectionWithDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetStaticSiteBuildDatabaseConnectionWithDetails.json
     */
    /**
     * Sample code: Get details of database connections for the static site build.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getDetailsOfDatabaseConnectionsForTheStaticSiteBuild(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().getBuildDatabaseConnectionWithDetailsWithResponse("rg",
            "testStaticSite0", "default", "default", com.azure.core.util.Context.NONE);
    }
}
