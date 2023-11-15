/** Samples for PacketCoreControlPlanes Reinstall. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-09-01/examples/PacketCoreControlPlaneReinstall.json
     */
    /**
     * Sample code: Reinstall packet core control plane.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void reinstallPacketCoreControlPlane(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.packetCoreControlPlanes().reinstall("rg1", "TestPacketCoreCP", com.azure.core.util.Context.NONE);
    }
}
