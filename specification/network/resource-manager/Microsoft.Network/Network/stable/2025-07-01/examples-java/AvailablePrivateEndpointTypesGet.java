
/**
 * Samples for AvailablePrivateEndpointTypes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/AvailablePrivateEndpointTypesGet.json
     */
    /**
     * Sample code: Get available PrivateEndpoint types.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getAvailablePrivateEndpointTypes(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAvailablePrivateEndpointTypes().list("regionName", com.azure.core.util.Context.NONE);
    }
}
