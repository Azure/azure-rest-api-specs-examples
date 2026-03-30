
/**
 * Samples for StaticSites GetLinkedBackends.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetLinkedBackendsForStaticSite.json
     */
    /**
     * Sample code: Get details of the linked backends registered with a static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getDetailsOfTheLinkedBackendsRegisteredWithAStaticSite(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().getLinkedBackends("rg", "testStaticSite0",
            com.azure.core.util.Context.NONE);
    }
}
