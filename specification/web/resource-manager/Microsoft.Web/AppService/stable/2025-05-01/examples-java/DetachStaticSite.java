
/**
 * Samples for StaticSites DetachStaticSite.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/DetachStaticSite.json
     */
    /**
     * Sample code: Detach a static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void detachAStaticSite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().detachStaticSite("rg", "testStaticSite0",
            com.azure.core.util.Context.NONE);
    }
}
