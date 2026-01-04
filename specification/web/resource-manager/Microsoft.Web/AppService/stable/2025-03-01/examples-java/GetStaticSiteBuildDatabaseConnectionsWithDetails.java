
/**
 * Samples for StaticSites GetBuildDatabaseConnectionsWithDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * GetStaticSiteBuildDatabaseConnectionsWithDetails.json
     */
    /**
     * Sample code: List full details of database connections for the static site build.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listFullDetailsOfDatabaseConnectionsForTheStaticSiteBuild(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().getBuildDatabaseConnectionsWithDetails("rg",
            "testStaticSite0", "default", com.azure.core.util.Context.NONE);
    }
}
