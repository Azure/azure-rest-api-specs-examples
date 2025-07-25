
import com.azure.resourcemanager.network.models.VpnGatewayPacketCaptureStopParameters;

/**
 * Samples for VpnGateways StopPacketCapture.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/VpnGatewayStopPacketCapture.
     * json
     */
    /**
     * Sample code: Stop packet capture on vpn gateway.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void stopPacketCaptureOnVpnGateway(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVpnGateways().stopPacketCapture("rg1", "vpngw",
            new VpnGatewayPacketCaptureStopParameters().withSasUrl(
                "https://teststorage.blob.core.windows.net/?sv=2018-03-28&ss=bfqt&srt=sco&sp=rwdlacup&se=2019-09-13T07:44:05Z&st=2019-09-06T23:44:05Z&spr=https&sig=V1h9D1riltvZMI69d6ihENnFo%2FrCvTqGgjO2lf%2FVBhE%3D"),
            com.azure.core.util.Context.NONE);
    }
}
