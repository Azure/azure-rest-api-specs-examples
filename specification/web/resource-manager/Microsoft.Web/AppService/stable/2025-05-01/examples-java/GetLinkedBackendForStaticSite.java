
/**
 * Samples for StaticSites GetLinkedBackend.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetLinkedBackendForStaticSite.json
     */
    /**
     * Sample code: Get details of the linked backend registered with a static site by name.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getDetailsOfTheLinkedBackendRegisteredWithAStaticSiteByName(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().getLinkedBackendWithResponse("rg", "testStaticSite0", "testBackend",
            com.azure.core.util.Context.NONE);
    }
}
