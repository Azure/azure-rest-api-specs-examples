
/**
 * Samples for PrivateAccesses ListPrivateEndpointConnections.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/PrivateAccesses_ListPrivateEndpointConnections.json
     */
    /**
     * Sample code: List all private endpoint connections under a private access resource.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void listAllPrivateEndpointConnectionsUnderAPrivateAccessResource(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.privateAccesses().listPrivateEndpointConnections("myResourceGroup", "myPrivateAccess",
            com.azure.core.util.Context.NONE);
    }
}
