
/**
 * Samples for PrivateEndpointConnections ListByFileShare.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/PrivateEndpointConnections_ListByFileShare.json
     */
    /**
     * Sample code: PrivateEndpointConnections_ListByFileShare.
     * 
     * @param manager Entry point to FileSharesManager.
     */
    public static void
        privateEndpointConnectionsListByFileShare(com.azure.resourcemanager.fileshares.FileSharesManager manager) {
        manager.privateEndpointConnections().listByFileShare("rgfileshares", "fileshare",
            com.azure.core.util.Context.NONE);
    }
}
