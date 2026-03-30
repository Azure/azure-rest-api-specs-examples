
/**
 * Samples for StaticSites GetUserProvidedFunctionAppForStaticSite.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetUserProvidedFunctionAppForStaticSite.json
     */
    /**
     * Sample code: Get details of the user provided function app registered with a static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getDetailsOfTheUserProvidedFunctionAppRegisteredWithAStaticSite(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().getUserProvidedFunctionAppForStaticSiteWithResponse("rg",
            "testStaticSite0", "testFunctionApp", com.azure.core.util.Context.NONE);
    }
}
