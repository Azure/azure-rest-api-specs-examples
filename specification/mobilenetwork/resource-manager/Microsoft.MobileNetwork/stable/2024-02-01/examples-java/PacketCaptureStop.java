
/**
 * Samples for PacketCaptures Stop.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/PacketCaptureStop
     * .json
     */
    /**
     * Sample code: Stop packet capture session.
     * 
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void stopPacketCaptureSession(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.packetCaptures().stop("rg1", "TestPacketCoreCP", "pc1", com.azure.core.util.Context.NONE);
    }
}
