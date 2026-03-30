
/**
 * Samples for WebApps GetPrivateEndpointConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetSitePrivateEndpointConnection.json
     */
    /**
     * Sample code: Get a private endpoint connection for a site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        getAPrivateEndpointConnectionForASite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getPrivateEndpointConnectionWithResponse("rg", "testSite", "connection",
            com.azure.core.util.Context.NONE);
    }
}
