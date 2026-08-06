
/**
 * Samples for BookshelfPrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/BookshelfPrivateEndpointConnections_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: BookshelfPrivateEndpointConnections_Delete_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void bookshelfPrivateEndpointConnectionsDeleteMaximumSet(
        com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.bookshelfPrivateEndpointConnections().delete("rgdiscovery", "f26e3436689dc08264", "connection",
            com.azure.core.util.Context.NONE);
    }
}
