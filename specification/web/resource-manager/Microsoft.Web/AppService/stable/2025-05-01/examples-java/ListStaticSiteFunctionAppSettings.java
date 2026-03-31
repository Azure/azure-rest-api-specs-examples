
/**
 * Samples for StaticSites ListStaticSiteFunctionAppSettings.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListStaticSiteFunctionAppSettings.json
     */
    /**
     * Sample code: Get function app settings of a static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        getFunctionAppSettingsOfAStaticSite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().listStaticSiteFunctionAppSettingsWithResponse("rg", "testStaticSite0",
            com.azure.core.util.Context.NONE);
    }
}
