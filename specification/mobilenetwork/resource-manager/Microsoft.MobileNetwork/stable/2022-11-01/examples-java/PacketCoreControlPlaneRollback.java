import com.azure.core.util.Context;

/** Samples for PacketCoreControlPlaneOperation Rollback. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/PacketCoreControlPlaneRollback.json
     */
    /**
     * Sample code: Rollback packet core control plane.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void rollbackPacketCoreControlPlane(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.packetCoreControlPlaneOperations().rollback("rg1", "TestPacketCoreCP", Context.NONE);
    }
}
