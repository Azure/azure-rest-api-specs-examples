/** Samples for PacketCoreControlPlanes Rollback. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-09-01/examples/PacketCoreControlPlaneRollback.json
     */
    /**
     * Sample code: Rollback packet core control plane.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void rollbackPacketCoreControlPlane(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.packetCoreControlPlanes().rollback("rg1", "TestPacketCoreCP", com.azure.core.util.Context.NONE);
    }
}
