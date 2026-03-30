
/**
 * Samples for StaticSites UnlinkBackend.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/UnlinkBackendFromStaticSite.json
     */
    /**
     * Sample code: Unlink a backend from a static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void unlinkABackendFromAStaticSite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().unlinkBackendWithResponse("rg", "testStaticSite0", "testBackend", null,
            com.azure.core.util.Context.NONE);
    }
}
