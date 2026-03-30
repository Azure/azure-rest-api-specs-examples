
import com.azure.resourcemanager.appservice.models.DatabaseConnectionPatchRequest;

/**
 * Samples for StaticSites UpdateDatabaseConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/PatchStaticSiteDatabaseConnection.json
     */
    /**
     * Sample code: Patch a database connection for a static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        patchADatabaseConnectionForAStaticSite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().updateDatabaseConnectionWithResponse("rg", "testStaticSite0",
            "default", new DatabaseConnectionPatchRequest(), com.azure.core.util.Context.NONE);
    }
}
