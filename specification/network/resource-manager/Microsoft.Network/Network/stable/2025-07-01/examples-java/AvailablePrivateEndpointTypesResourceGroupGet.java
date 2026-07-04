
/**
 * Samples for AvailablePrivateEndpointTypes ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/AvailablePrivateEndpointTypesResourceGroupGet.json
     */
    /**
     * Sample code: Get available PrivateEndpoint types in the resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getAvailablePrivateEndpointTypesInTheResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAvailablePrivateEndpointTypes().listByResourceGroup("regionName", "rg1",
            com.azure.core.util.Context.NONE);
    }
}
