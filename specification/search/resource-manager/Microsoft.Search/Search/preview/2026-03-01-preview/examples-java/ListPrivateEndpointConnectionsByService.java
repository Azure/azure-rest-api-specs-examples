
/**
 * Samples for PrivateEndpointConnections ListByService.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ListPrivateEndpointConnectionsByService.json
     */
    /**
     * Sample code: ListPrivateEndpointConnectionsByService.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void
        listPrivateEndpointConnectionsByService(com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().listByService("rg1", "mysearchservice",
            com.azure.core.util.Context.NONE);
    }
}
