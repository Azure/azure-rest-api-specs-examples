
/**
 * Samples for StaticSites Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/DeleteStaticSite.json
     */
    /**
     * Sample code: Delete a static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void deleteAStaticSite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().delete("rg", "testStaticSite0", com.azure.core.util.Context.NONE);
    }
}
