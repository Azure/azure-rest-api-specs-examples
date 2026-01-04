
import com.azure.resourcemanager.appservice.models.DatabaseConnectionPatchRequest;

/**
 * Samples for StaticSites UpdateBuildDatabaseConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * PatchStaticSiteBuildDatabaseConnection.json
     */
    /**
     * Sample code: Patch a database connection for a static site build.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        patchADatabaseConnectionForAStaticSiteBuild(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().updateBuildDatabaseConnectionWithResponse("rg",
            "testStaticSite0", "default", "default", new DatabaseConnectionPatchRequest(),
            com.azure.core.util.Context.NONE);
    }
}
