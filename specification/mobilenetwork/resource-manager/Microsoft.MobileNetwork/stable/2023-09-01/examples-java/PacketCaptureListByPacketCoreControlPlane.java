/** Samples for PacketCaptures ListByPacketCoreControlPlane. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-09-01/examples/PacketCaptureListByPacketCoreControlPlane.json
     */
    /**
     * Sample code: List packet capture sessions under a packet core control plane.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listPacketCaptureSessionsUnderAPacketCoreControlPlane(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .packetCaptures()
            .listByPacketCoreControlPlane("rg1", "TestPacketCoreCP", com.azure.core.util.Context.NONE);
    }
}
