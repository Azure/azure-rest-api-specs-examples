
/**
 * Samples for StaticSites GetUserProvidedFunctionAppForStaticSiteBuild.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetUserProvidedFunctionAppForStaticSiteBuild.json
     */
    /**
     * Sample code: Get details of the user provided function app registered with a static site build.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getDetailsOfTheUserProvidedFunctionAppRegisteredWithAStaticSiteBuild(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().getUserProvidedFunctionAppForStaticSiteBuildWithResponse("rg",
            "testStaticSite0", "default", "testFunctionApp", com.azure.core.util.Context.NONE);
    }
}
