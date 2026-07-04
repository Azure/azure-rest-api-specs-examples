
import com.azure.resourcemanager.network.models.VpnPacketCaptureStartParameters;

/**
 * Samples for VirtualNetworkGateways StartPacketCapture.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayStartPacketCaptureFilterData.json
     */
    /**
     * Sample code: Start packet capture on virtual network gateway with filter.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        startPacketCaptureOnVirtualNetworkGatewayWithFilter(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().startPacketCapture("rg1", "vpngw",
            new VpnPacketCaptureStartParameters().withFilterData(
                "{'TracingFlags': 11,'MaxPacketBufferSize': 120,'MaxFileSize': 200,'Filters': [{'SourceSubnets': ['20.1.1.0/24'],'DestinationSubnets': ['10.1.1.0/24'],'SourcePort': [500],'DestinationPort': [4500],'Protocol': 6,'TcpFlags': 16,'CaptureSingleDirectionTrafficOnly': true}]}"),
            com.azure.core.util.Context.NONE);
    }
}
