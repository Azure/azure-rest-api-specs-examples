
/**
 * Samples for WebApps GetPrivateEndpointConnectionSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetSitePrivateEndpointConnectionSlot.json
     */
    /**
     * Sample code: Get a private endpoint connection for a site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        getAPrivateEndpointConnectionForASite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getPrivateEndpointConnectionSlotWithResponse("rg", "testSite",
            "connection", "stage", com.azure.core.util.Context.NONE);
    }
}
