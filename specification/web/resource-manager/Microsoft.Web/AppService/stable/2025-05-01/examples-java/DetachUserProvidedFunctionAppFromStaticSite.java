
/**
 * Samples for StaticSites DetachUserProvidedFunctionAppFromStaticSite.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/DetachUserProvidedFunctionAppFromStaticSite.json
     */
    /**
     * Sample code: Detach the user provided function app from the static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void detachTheUserProvidedFunctionAppFromTheStaticSite(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().detachUserProvidedFunctionAppFromStaticSiteWithResponse("rg",
            "testStaticSite0", "testFunctionApp", com.azure.core.util.Context.NONE);
    }
}
