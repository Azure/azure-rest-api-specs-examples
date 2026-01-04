
import com.azure.resourcemanager.network.models.VpnClientParameters;

/**
 * Samples for VirtualNetworkGateways Generatevpnclientpackage.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * VirtualNetworkGatewayGenerateVpnClientPackage.json
     */
    /**
     * Sample code: GenerateVPNClientPackage.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void generateVPNClientPackage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGateways().generatevpnclientpackage("rg1", "vpngw",
            new VpnClientParameters(), com.azure.core.util.Context.NONE);
    }
}
