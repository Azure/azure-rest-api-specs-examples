
/**
 * Samples for WebApps GetPrivateLinkResourcesSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetSitePrivateLinkResourcesSlot.json
     */
    /**
     * Sample code: Get private link resources of a site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getPrivateLinkResourcesOfASite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getPrivateLinkResourcesSlotWithResponse("rg", "testSite", "stage",
            com.azure.core.util.Context.NONE);
    }
}
