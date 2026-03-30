
/**
 * Samples for StaticSites ListStaticSiteCustomDomains.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetStaticSiteCustomDomains.json
     */
    /**
     * Sample code: List custom domains for a static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listCustomDomainsForAStaticSite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().listStaticSiteCustomDomains("rg", "testStaticSite0",
            com.azure.core.util.Context.NONE);
    }
}
