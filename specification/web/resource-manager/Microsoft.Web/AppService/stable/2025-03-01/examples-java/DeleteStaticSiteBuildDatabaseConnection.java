
/**
 * Samples for StaticSites DeleteBuildDatabaseConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * DeleteStaticSiteBuildDatabaseConnection.json
     */
    /**
     * Sample code: Delete a database connection from a static site build.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        deleteADatabaseConnectionFromAStaticSiteBuild(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().deleteBuildDatabaseConnectionWithResponse("rg",
            "testStaticSite0", "default", "default", com.azure.core.util.Context.NONE);
    }
}
