/** Samples for PacketCoreDataPlanes Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-06-01/examples/PacketCoreDataPlaneDelete.json
     */
    /**
     * Sample code: Delete packet core data plane.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void deletePacketCoreDataPlane(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .packetCoreDataPlanes()
            .delete("rg1", "testPacketCoreCP", "testPacketCoreDP", com.azure.core.util.Context.NONE);
    }
}
