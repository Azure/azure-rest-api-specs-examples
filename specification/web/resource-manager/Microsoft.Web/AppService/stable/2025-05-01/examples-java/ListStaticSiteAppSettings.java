
/**
 * Samples for StaticSites ListStaticSiteAppSettings.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListStaticSiteAppSettings.json
     */
    /**
     * Sample code: Get app settings of a static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAppSettingsOfAStaticSite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().listStaticSiteAppSettingsWithResponse("rg", "testStaticSite0",
            com.azure.core.util.Context.NONE);
    }
}
