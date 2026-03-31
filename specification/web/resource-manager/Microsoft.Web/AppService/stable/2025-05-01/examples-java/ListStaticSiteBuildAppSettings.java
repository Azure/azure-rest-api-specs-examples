
/**
 * Samples for StaticSites ListStaticSiteBuildAppSettings.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListStaticSiteBuildAppSettings.json
     */
    /**
     * Sample code: Get app settings of a static site build.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        getAppSettingsOfAStaticSiteBuild(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().listStaticSiteBuildAppSettingsWithResponse("rg", "testStaticSite0",
            "12", com.azure.core.util.Context.NONE);
    }
}
