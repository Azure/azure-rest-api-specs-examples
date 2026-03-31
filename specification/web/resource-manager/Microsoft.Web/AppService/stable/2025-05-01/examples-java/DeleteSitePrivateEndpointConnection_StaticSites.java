
/**
 * Samples for StaticSites DeletePrivateEndpointConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/DeleteSitePrivateEndpointConnection_StaticSites.json
     */
    /**
     * Sample code: Delete a private endpoint connection for a site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        deleteAPrivateEndpointConnectionForASite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().deletePrivateEndpointConnection("rg", "testSite", "connection",
            com.azure.core.util.Context.NONE);
    }
}
