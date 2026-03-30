
import com.azure.resourcemanager.appservice.models.DatabaseConnectionPatchRequest;

/**
 * Samples for StaticSites UpdateBuildDatabaseConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/PatchStaticSiteBuildDatabaseConnection.json
     */
    /**
     * Sample code: Patch a database connection for a static site build.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        patchADatabaseConnectionForAStaticSiteBuild(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().updateBuildDatabaseConnectionWithResponse("rg", "testStaticSite0",
            "default", "default", new DatabaseConnectionPatchRequest(), com.azure.core.util.Context.NONE);
    }
}
