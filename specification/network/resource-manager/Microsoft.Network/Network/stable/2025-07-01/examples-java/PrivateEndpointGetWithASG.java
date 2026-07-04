
/**
 * Samples for PrivateEndpoints GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PrivateEndpointGetWithASG.json
     */
    /**
     * Sample code: Get private endpoint with application security groups.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getPrivateEndpointWithApplicationSecurityGroups(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPrivateEndpoints().getByResourceGroupWithResponse("rg1", "testPe", null,
            com.azure.core.util.Context.NONE);
    }
}
