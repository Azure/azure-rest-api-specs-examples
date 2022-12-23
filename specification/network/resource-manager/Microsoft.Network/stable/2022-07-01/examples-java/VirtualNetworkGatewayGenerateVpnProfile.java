import com.azure.core.util.Context;
import com.azure.resourcemanager.network.models.VpnClientParameters;

/** Samples for VirtualNetworkGateways GenerateVpnProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/VirtualNetworkGatewayGenerateVpnProfile.json
     */
    /**
     * Sample code: GenerateVirtualNetworkGatewayVPNProfile.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void generateVirtualNetworkGatewayVPNProfile(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualNetworkGateways()
            .generateVpnProfile("rg1", "vpngw", new VpnClientParameters(), Context.NONE);
    }
}
