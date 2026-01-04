
/**
 * Samples for StaticSites DeleteDatabaseConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * DeleteStaticSiteDatabaseConnection.json
     */
    /**
     * Sample code: Delete a database connection from a static site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteADatabaseConnectionFromAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().deleteDatabaseConnectionWithResponse("rg",
            "testStaticSite0", "default", com.azure.core.util.Context.NONE);
    }
}
