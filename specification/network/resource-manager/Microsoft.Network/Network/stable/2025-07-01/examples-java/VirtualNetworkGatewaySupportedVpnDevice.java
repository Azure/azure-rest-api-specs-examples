
/**
 * Samples for VirtualNetworkGateways SupportedVpnDevices.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewaySupportedVpnDevice.json
     */
    /**
     * Sample code: ListVirtualNetworkGatewaySupportedVPNDevices.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listVirtualNetworkGatewaySupportedVPNDevices(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().supportedVpnDevicesWithResponse("rg1", "vpngw",
            com.azure.core.util.Context.NONE);
    }
}
