
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/DeletePrivateEndpointConnection.json
     */
    /**
     * Sample code: PrivateEndpointConnectionDelete.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void privateEndpointConnectionDelete(com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().deleteWithResponse("rg1", "mysearchservice",
            "testEndpoint.50bf4fbe-d7c1-4b48-a642-4f5892642546", com.azure.core.util.Context.NONE);
    }
}
