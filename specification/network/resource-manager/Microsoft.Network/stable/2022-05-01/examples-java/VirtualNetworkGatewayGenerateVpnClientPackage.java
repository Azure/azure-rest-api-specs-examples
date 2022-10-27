import com.azure.core.util.Context;
import com.azure.resourcemanager.network.models.VpnClientParameters;

/** Samples for VirtualNetworkGateways Generatevpnclientpackage. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/VirtualNetworkGatewayGenerateVpnClientPackage.json
     */
    /**
     * Sample code: GenerateVPNClientPackage.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void generateVPNClientPackage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualNetworkGateways()
            .generatevpnclientpackage("rg1", "vpngw", new VpnClientParameters(), Context.NONE);
    }
}
