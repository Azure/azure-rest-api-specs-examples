
/**
 * Samples for PrivateEndpoints ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PrivateEndpointList.json
     */
    /**
     * Sample code: List private endpoints in resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listPrivateEndpointsInResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPrivateEndpoints().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
