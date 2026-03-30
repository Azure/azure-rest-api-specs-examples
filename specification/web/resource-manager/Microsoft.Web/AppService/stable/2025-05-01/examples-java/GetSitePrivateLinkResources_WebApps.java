
/**
 * Samples for WebApps GetPrivateLinkResources.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetSitePrivateLinkResources_WebApps.json
     */
    /**
     * Sample code: Get private link resources of a site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getPrivateLinkResourcesOfASite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getPrivateLinkResourcesWithResponse("rg", "testSite",
            com.azure.core.util.Context.NONE);
    }
}
