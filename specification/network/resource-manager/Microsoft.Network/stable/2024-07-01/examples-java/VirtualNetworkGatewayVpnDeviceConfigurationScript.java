
import com.azure.resourcemanager.network.models.VpnDeviceScriptParameters;

/**
 * Samples for VirtualNetworkGateways VpnDeviceConfigurationScript.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/
     * VirtualNetworkGatewayVpnDeviceConfigurationScript.json
     */
    /**
     * Sample code: GetVPNDeviceConfigurationScript.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVPNDeviceConfigurationScript(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGateways()
            .vpnDeviceConfigurationScriptWithResponse("rg1", "vpngw", new VpnDeviceScriptParameters()
                .withVendor("Cisco").withDeviceFamily("ISR").withFirmwareVersion("IOS 15.1 (Preview)"),
                com.azure.core.util.Context.NONE);
    }
}
