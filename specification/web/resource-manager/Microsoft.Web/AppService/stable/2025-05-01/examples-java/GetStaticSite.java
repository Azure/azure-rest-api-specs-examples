
/**
 * Samples for StaticSites GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetStaticSite.json
     */
    /**
     * Sample code: Get details for a static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getDetailsForAStaticSite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().getByResourceGroupWithResponse("rg", "testStaticSite0",
            com.azure.core.util.Context.NONE);
    }
}
