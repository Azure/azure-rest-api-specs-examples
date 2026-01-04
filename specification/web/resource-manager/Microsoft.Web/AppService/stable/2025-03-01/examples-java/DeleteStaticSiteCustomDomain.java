
/**
 * Samples for StaticSites DeleteStaticSiteCustomDomain.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * DeleteStaticSiteCustomDomain.json
     */
    /**
     * Sample code: Delete a custom domain for a static site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteACustomDomainForAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().deleteStaticSiteCustomDomain("rg", "testStaticSite0",
            "custom.domain.net", com.azure.core.util.Context.NONE);
    }
}
