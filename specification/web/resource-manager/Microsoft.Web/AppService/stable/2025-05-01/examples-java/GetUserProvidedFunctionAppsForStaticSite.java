
/**
 * Samples for StaticSites GetUserProvidedFunctionAppsForStaticSite.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetUserProvidedFunctionAppsForStaticSite.json
     */
    /**
     * Sample code: Get details of the user provided function apps registered with a static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getDetailsOfTheUserProvidedFunctionAppsRegisteredWithAStaticSite(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().getUserProvidedFunctionAppsForStaticSite("rg", "testStaticSite0",
            com.azure.core.util.Context.NONE);
    }
}
