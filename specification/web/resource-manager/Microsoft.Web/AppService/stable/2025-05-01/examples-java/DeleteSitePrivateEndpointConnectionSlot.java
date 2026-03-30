
/**
 * Samples for WebApps DeletePrivateEndpointConnectionSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/DeleteSitePrivateEndpointConnectionSlot.json
     */
    /**
     * Sample code: Delete a private endpoint connection for a site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        deleteAPrivateEndpointConnectionForASite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().deletePrivateEndpointConnectionSlot("rg", "testSite", "connection",
            "stage", com.azure.core.util.Context.NONE);
    }
}
