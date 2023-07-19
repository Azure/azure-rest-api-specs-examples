/** Samples for PacketCoreDataPlanes Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-06-01/examples/PacketCoreDataPlaneGet.json
     */
    /**
     * Sample code: Get packet core data plane.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getPacketCoreDataPlane(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .packetCoreDataPlanes()
            .getWithResponse("rg1", "testPacketCoreCP", "testPacketCoreDP", com.azure.core.util.Context.NONE);
    }
}
