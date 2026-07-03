
import com.azure.resourcemanager.network.models.P2SVpnConnectionRequest;
import java.util.Arrays;

/**
 * Samples for VirtualNetworkGateways DisconnectVirtualNetworkGatewayVpnConnections.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewaysDisconnectP2sVpnConnections.json
     */
    /**
     * Sample code: Disconnect VpnConnections from Virtual Network Gateway.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        disconnectVpnConnectionsFromVirtualNetworkGateway(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().disconnectVirtualNetworkGatewayVpnConnections(
            "vpn-gateway-test", "vpngateway",
            new P2SVpnConnectionRequest().withVpnConnectionIds(Arrays.asList("vpnconnId1", "vpnconnId2")),
            com.azure.core.util.Context.NONE);
    }
}
