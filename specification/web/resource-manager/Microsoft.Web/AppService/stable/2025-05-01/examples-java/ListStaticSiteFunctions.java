
/**
 * Samples for StaticSites ListStaticSiteFunctions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListStaticSiteFunctions.json
     */
    /**
     * Sample code: Gets the functions of a static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getsTheFunctionsOfAStaticSite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().listStaticSiteFunctions("rg", "testStaticSite0",
            com.azure.core.util.Context.NONE);
    }
}
