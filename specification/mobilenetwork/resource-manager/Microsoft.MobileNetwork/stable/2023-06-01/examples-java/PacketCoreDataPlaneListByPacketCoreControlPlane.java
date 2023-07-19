/** Samples for PacketCoreDataPlanes ListByPacketCoreControlPlane. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-06-01/examples/PacketCoreDataPlaneListByPacketCoreControlPlane.json
     */
    /**
     * Sample code: List packet core data planes in a control plane.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listPacketCoreDataPlanesInAControlPlane(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .packetCoreDataPlanes()
            .listByPacketCoreControlPlane("rg1", "testPacketCoreCP", com.azure.core.util.Context.NONE);
    }
}
