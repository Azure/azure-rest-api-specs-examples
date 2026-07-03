
import com.azure.resourcemanager.network.models.VpnPacketCaptureStartParameters;

/**
 * Samples for VirtualNetworkGatewayConnections StartPacketCapture.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayConnectionStartPacketCaptureFilterData.json
     */
    /**
     * Sample code: Start packet capture on virtual network gateway connection with filter.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void startPacketCaptureOnVirtualNetworkGatewayConnectionWithFilter(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGatewayConnections().startPacketCapture("rg1", "vpngwcn1",
            new VpnPacketCaptureStartParameters().withFilterData(
                "{'TracingFlags': 11,'MaxPacketBufferSize': 120,'MaxFileSize': 200,'Filters': [{'SourceSubnets': ['20.1.1.0/24'],'DestinationSubnets': ['10.1.1.0/24'],'SourcePort': [500],'DestinationPort': [4500],'Protocol': 6,'TcpFlags': 16,'CaptureSingleDirectionTrafficOnly': true}]}"),
            com.azure.core.util.Context.NONE);
    }
}
