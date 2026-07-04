
/**
 * Samples for LocalNetworkGateways Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/LocalNetworkGatewayDelete.json
     */
    /**
     * Sample code: DeleteLocalNetworkGateway.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteLocalNetworkGateway(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getLocalNetworkGateways().delete("rg1", "localgw", com.azure.core.util.Context.NONE);
    }
}
