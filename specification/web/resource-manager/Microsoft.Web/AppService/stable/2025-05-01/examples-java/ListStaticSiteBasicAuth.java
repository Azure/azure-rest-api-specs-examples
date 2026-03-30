
/**
 * Samples for StaticSites ListBasicAuth.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListStaticSiteBasicAuth.json
     */
    /**
     * Sample code: Lists the basic auth properties for a static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        listsTheBasicAuthPropertiesForAStaticSite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().listBasicAuth("rg", "testStaticSite0",
            com.azure.core.util.Context.NONE);
    }
}
