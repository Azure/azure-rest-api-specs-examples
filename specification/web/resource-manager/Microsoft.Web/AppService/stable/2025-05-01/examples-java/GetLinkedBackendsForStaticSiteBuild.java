
/**
 * Samples for StaticSites GetLinkedBackendsForBuild.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetLinkedBackendsForStaticSiteBuild.json
     */
    /**
     * Sample code: Get details of the linked backends registered with a static site build.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getDetailsOfTheLinkedBackendsRegisteredWithAStaticSiteBuild(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().getLinkedBackendsForBuild("rg", "testStaticSite0", "default",
            com.azure.core.util.Context.NONE);
    }
}
