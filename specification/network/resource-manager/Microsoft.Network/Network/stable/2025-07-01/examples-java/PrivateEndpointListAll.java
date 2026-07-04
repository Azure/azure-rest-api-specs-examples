
/**
 * Samples for PrivateEndpoints List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PrivateEndpointListAll.json
     */
    /**
     * Sample code: List all private endpoints.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllPrivateEndpoints(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPrivateEndpoints().list(com.azure.core.util.Context.NONE);
    }
}
