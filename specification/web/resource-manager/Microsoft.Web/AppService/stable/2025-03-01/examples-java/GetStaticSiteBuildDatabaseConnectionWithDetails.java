
/**
 * Samples for StaticSites GetBuildDatabaseConnectionWithDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * GetStaticSiteBuildDatabaseConnectionWithDetails.json
     */
    /**
     * Sample code: Get details of database connections for the static site build.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getDetailsOfDatabaseConnectionsForTheStaticSiteBuild(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().getBuildDatabaseConnectionWithDetailsWithResponse(
            "rg", "testStaticSite0", "default", "default", com.azure.core.util.Context.NONE);
    }
}
