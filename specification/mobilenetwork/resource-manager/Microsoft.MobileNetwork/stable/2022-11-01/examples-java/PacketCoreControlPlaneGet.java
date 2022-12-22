import com.azure.core.util.Context;

/** Samples for PacketCoreControlPlanes GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/PacketCoreControlPlaneGet.json
     */
    /**
     * Sample code: Get packet core control plane.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getPacketCoreControlPlane(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.packetCoreControlPlanes().getByResourceGroupWithResponse("rg1", "TestPacketCoreCP", Context.NONE);
    }
}
