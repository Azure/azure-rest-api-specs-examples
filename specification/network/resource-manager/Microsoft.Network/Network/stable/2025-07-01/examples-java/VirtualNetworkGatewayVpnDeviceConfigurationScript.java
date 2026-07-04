
import com.azure.resourcemanager.network.models.VpnDeviceScriptParameters;

/**
 * Samples for VirtualNetworkGateways VpnDeviceConfigurationScript.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayVpnDeviceConfigurationScript.json
     */
    /**
     * Sample code: GetVPNDeviceConfigurationScript.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getVPNDeviceConfigurationScript(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways()
            .vpnDeviceConfigurationScriptWithResponse("rg1", "vpngw", new VpnDeviceScriptParameters()
                .withVendor("Cisco").withDeviceFamily("ISR").withFirmwareVersion("IOS 15.1 (Preview)"),
                com.azure.core.util.Context.NONE);
    }
}
