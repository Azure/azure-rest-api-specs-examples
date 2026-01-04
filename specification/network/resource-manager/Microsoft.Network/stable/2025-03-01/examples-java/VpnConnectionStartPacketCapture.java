
import com.azure.resourcemanager.network.models.VpnConnectionPacketCaptureStartParameters;
import java.util.Arrays;

/**
 * Samples for VpnConnections StartPacketCapture.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * VpnConnectionStartPacketCapture.json
     */
    /**
     * Sample code: Start packet capture on vpn connection without filter.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        startPacketCaptureOnVpnConnectionWithoutFilter(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVpnConnections()
            .startPacketCapture(
                "rg1", "gateway1", "vpnConnection1", new VpnConnectionPacketCaptureStartParameters()
                    .withLinkConnectionNames(Arrays.asList("siteLink1", "siteLink2")),
                com.azure.core.util.Context.NONE);
    }
}
