import com.azure.core.util.Context;

/** Samples for VirtualNetworkGateways ListConnections. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VirtualNetworkGatewaysListConnections.json
     */
    /**
     * Sample code: VirtualNetworkGatewaysListConnections.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualNetworkGatewaysListConnections(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualNetworkGateways()
            .listConnections("testrg", "test-vpn-gateway-1", Context.NONE);
    }
}
