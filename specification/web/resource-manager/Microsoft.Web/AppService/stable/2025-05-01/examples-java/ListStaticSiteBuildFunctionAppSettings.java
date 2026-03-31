
/**
 * Samples for StaticSites ListStaticSiteBuildFunctionAppSettings.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListStaticSiteBuildFunctionAppSettings.json
     */
    /**
     * Sample code: Get function app settings of a static site build.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        getFunctionAppSettingsOfAStaticSiteBuild(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().listStaticSiteBuildFunctionAppSettingsWithResponse("rg",
            "testStaticSite0", "12", com.azure.core.util.Context.NONE);
    }
}
