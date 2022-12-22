import com.azure.core.util.Context;

/** Samples for PacketCoreControlPlanes Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/PacketCoreControlPlaneDelete.json
     */
    /**
     * Sample code: Delete packet core control plane.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void deletePacketCoreControlPlane(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.packetCoreControlPlanes().delete("rg1", "TestPacketCoreCP", Context.NONE);
    }
}
