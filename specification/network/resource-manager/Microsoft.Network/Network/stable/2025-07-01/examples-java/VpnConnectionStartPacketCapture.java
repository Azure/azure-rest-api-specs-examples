
import com.azure.resourcemanager.network.models.VpnConnectionPacketCaptureStartParameters;
import java.util.Arrays;

/**
 * Samples for VpnConnections StartPacketCapture.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnConnectionStartPacketCapture.json
     */
    /**
     * Sample code: Start packet capture on vpn connection without filter.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        startPacketCaptureOnVpnConnectionWithoutFilter(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnConnections()
            .startPacketCapture(
                "rg1", "gateway1", "vpnConnection1", new VpnConnectionPacketCaptureStartParameters()
                    .withLinkConnectionNames(Arrays.asList("siteLink1", "siteLink2")),
                com.azure.core.util.Context.NONE);
    }
}
