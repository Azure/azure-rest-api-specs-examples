
/**
 * Samples for StaticSites GetBuildDatabaseConnections.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * GetStaticSiteBuildDatabaseConnections.json
     */
    /**
     * Sample code: List overviews of database connections for the static site build.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listOverviewsOfDatabaseConnectionsForTheStaticSiteBuild(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().getBuildDatabaseConnections("rg", "testStaticSite0",
            "default", com.azure.core.util.Context.NONE);
    }
}
