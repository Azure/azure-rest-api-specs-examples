import com.azure.core.util.Context;

/** Samples for VirtualNetworkGateways SupportedVpnDevices. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VirtualNetworkGatewaySupportedVpnDevice.json
     */
    /**
     * Sample code: ListVirtualNetworkGatewaySupportedVPNDevices.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listVirtualNetworkGatewaySupportedVPNDevices(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualNetworkGateways()
            .supportedVpnDevicesWithResponse("rg1", "vpngw", Context.NONE);
    }
}
