
/**
 * Samples for StaticSites GetPrivateEndpointConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetSitePrivateEndpointConnection_StaticSites.json
     */
    /**
     * Sample code: Get a private endpoint connection for a site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        getAPrivateEndpointConnectionForASite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().getPrivateEndpointConnectionWithResponse("rg", "testSite",
            "connection", com.azure.core.util.Context.NONE);
    }
}
