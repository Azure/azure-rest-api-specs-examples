
/**
 * Samples for PrivateEndpoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/PrivateEndpointDelete.json
     */
    /**
     * Sample code: Delete private endpoint.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deletePrivateEndpoint(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPrivateEndpoints().delete("rg1", "testPe", com.azure.core.util.Context.NONE);
    }
}
