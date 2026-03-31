
/**
 * Samples for StaticSites GetUserProvidedFunctionAppsForStaticSiteBuild.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetUserProvidedFunctionAppsForStaticSiteBuild.json
     */
    /**
     * Sample code: Get details of the user provided function apps registered with a static site build.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getDetailsOfTheUserProvidedFunctionAppsRegisteredWithAStaticSiteBuild(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().getUserProvidedFunctionAppsForStaticSiteBuild("rg", "testStaticSite0",
            "default", com.azure.core.util.Context.NONE);
    }
}
