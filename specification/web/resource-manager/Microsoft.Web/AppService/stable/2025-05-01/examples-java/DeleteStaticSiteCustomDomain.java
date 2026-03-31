
/**
 * Samples for StaticSites DeleteStaticSiteCustomDomain.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/DeleteStaticSiteCustomDomain.json
     */
    /**
     * Sample code: Delete a custom domain for a static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        deleteACustomDomainForAStaticSite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().deleteStaticSiteCustomDomain("rg", "testStaticSite0",
            "custom.domain.net", com.azure.core.util.Context.NONE);
    }
}
