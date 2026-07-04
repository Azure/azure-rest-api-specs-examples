
/**
 * Samples for LocalNetworkGateways ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/LocalNetworkGatewayList.json
     */
    /**
     * Sample code: ListLocalNetworkGateways.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listLocalNetworkGateways(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getLocalNetworkGateways().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
