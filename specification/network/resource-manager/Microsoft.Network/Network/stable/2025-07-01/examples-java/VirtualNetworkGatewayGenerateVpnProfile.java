
import com.azure.resourcemanager.network.models.VpnClientParameters;

/**
 * Samples for VirtualNetworkGateways GenerateVpnProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayGenerateVpnProfile.json
     */
    /**
     * Sample code: GenerateVirtualNetworkGatewayVPNProfile.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        generateVirtualNetworkGatewayVPNProfile(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().generateVpnProfile("rg1", "vpngw",
            new VpnClientParameters(), com.azure.core.util.Context.NONE);
    }
}
