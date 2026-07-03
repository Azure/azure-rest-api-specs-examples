
import com.azure.resourcemanager.network.models.VpnClientParameters;

/**
 * Samples for VirtualNetworkGateways Generatevpnclientpackage.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayGenerateVpnClientPackage.json
     */
    /**
     * Sample code: GenerateVPNClientPackage.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void generateVPNClientPackage(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().generatevpnclientpackage("rg1", "vpngw",
            new VpnClientParameters(), com.azure.core.util.Context.NONE);
    }
}
