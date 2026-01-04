
import com.azure.resourcemanager.appservice.models.DatabaseConnectionPatchRequest;

/**
 * Samples for StaticSites UpdateDatabaseConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * PatchStaticSiteDatabaseConnection.json
     */
    /**
     * Sample code: Patch a database connection for a static site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void patchADatabaseConnectionForAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().updateDatabaseConnectionWithResponse("rg",
            "testStaticSite0", "default", new DatabaseConnectionPatchRequest(), com.azure.core.util.Context.NONE);
    }
}
