
/**
 * Samples for WebApps GetPrivateLinkResources.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * GetSitePrivateLinkResources.json
     */
    /**
     * Sample code: Get private link resources of a site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getPrivateLinkResourcesOfASite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().getPrivateLinkResourcesWithResponse("rg", "testSite",
            com.azure.core.util.Context.NONE);
    }
}
