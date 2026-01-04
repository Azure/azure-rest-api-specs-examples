
/**
 * Samples for StaticSites DeleteStaticSiteUser.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/DeleteStaticSiteUser.json
     */
    /**
     * Sample code: Delete a user for a static site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAUserForAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().deleteStaticSiteUserWithResponse("rg",
            "testStaticSite0", "aad", "1234", com.azure.core.util.Context.NONE);
    }
}
