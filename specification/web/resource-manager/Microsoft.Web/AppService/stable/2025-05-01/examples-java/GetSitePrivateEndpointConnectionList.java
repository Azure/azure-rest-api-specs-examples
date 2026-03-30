
/**
 * Samples for StaticSites GetPrivateEndpointConnectionList.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetSitePrivateEndpointConnectionList.json
     */
    /**
     * Sample code: Get a list of private endpoint connections associated with a site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAListOfPrivateEndpointConnectionsAssociatedWithASite(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().getPrivateEndpointConnectionList("rg", "testStaticSite0",
            com.azure.core.util.Context.NONE);
    }
}
