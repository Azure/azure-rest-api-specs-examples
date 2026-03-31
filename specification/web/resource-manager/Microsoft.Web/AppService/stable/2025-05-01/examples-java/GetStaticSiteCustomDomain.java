
/**
 * Samples for StaticSites GetStaticSiteCustomDomain.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetStaticSiteCustomDomain.json
     */
    /**
     * Sample code: Get custom domain for a static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getCustomDomainForAStaticSite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().getStaticSiteCustomDomainWithResponse("rg", "testStaticSite0",
            "custom.domain.net", com.azure.core.util.Context.NONE);
    }
}
