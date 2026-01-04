
/**
 * Samples for StaticSites GetStaticSiteCustomDomain.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/GetStaticSiteCustomDomain.
     * json
     */
    /**
     * Sample code: Get custom domain for a static site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getCustomDomainForAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().getStaticSiteCustomDomainWithResponse("rg",
            "testStaticSite0", "custom.domain.net", com.azure.core.util.Context.NONE);
    }
}
