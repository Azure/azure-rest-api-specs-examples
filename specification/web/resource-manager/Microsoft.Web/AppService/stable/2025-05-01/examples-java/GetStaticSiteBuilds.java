
/**
 * Samples for StaticSites GetStaticSiteBuilds.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetStaticSiteBuilds.json
     */
    /**
     * Sample code: Get all builds for a static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAllBuildsForAStaticSite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().getStaticSiteBuilds("rg", "testStaticSite0",
            com.azure.core.util.Context.NONE);
    }
}
