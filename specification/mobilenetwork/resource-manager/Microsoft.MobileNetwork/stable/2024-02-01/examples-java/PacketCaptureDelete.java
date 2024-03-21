
/**
 * Samples for PacketCaptures Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/
     * PacketCaptureDelete.json
     */
    /**
     * Sample code: Delete packet capture.
     * 
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void deletePacketCapture(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.packetCaptures().delete("rg1", "TestPacketCoreCP", "pc1", com.azure.core.util.Context.NONE);
    }
}
