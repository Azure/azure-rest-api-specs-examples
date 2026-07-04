
import com.azure.resourcemanager.network.models.VpnGatewayPacketCaptureStartParameters;

/**
 * Samples for VpnGateways StartPacketCapture.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnGatewayStartPacketCaptureFilterData.json
     */
    /**
     * Sample code: Start packet capture on vpn gateway with filter.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        startPacketCaptureOnVpnGatewayWithFilter(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnGateways().startPacketCapture("rg1", "vpngw",
            new VpnGatewayPacketCaptureStartParameters().withFilterData(
                "{'TracingFlags': 11,'MaxPacketBufferSize': 120,'MaxFileSize': 200,'Filters': [{'SourceSubnets': ['20.1.1.0/24'],'DestinationSubnets': ['10.1.1.0/24'],'SourcePort': [500],'DestinationPort': [4500],'Protocol': 6,'TcpFlags': 16,'CaptureSingleDirectionTrafficOnly': true}]}"),
            com.azure.core.util.Context.NONE);
    }
}
