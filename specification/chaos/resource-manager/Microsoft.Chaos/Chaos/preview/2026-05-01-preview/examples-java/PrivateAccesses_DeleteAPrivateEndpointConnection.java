
/**
 * Samples for PrivateAccesses DeleteAPrivateEndpointConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/PrivateAccesses_DeleteAPrivateEndpointConnection.json
     */
    /**
     * Sample code: Delete a private endpoint connection under a private access resource.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void deleteAPrivateEndpointConnectionUnderAPrivateAccessResource(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.privateAccesses().deleteAPrivateEndpointConnection("myResourceGroup", "myPrivateAccess",
            "myPrivateEndpointConnection", com.azure.core.util.Context.NONE);
    }
}
