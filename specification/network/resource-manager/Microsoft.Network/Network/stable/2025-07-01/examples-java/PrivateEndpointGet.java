
/**
 * Samples for PrivateEndpoints GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PrivateEndpointGet.json
     */
    /**
     * Sample code: Get private endpoint.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getPrivateEndpoint(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPrivateEndpoints().getByResourceGroupWithResponse("rg1", "testPe", null,
            com.azure.core.util.Context.NONE);
    }
}
