/** Samples for PacketCoreControlPlanes GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-06-01/examples/PacketCoreControlPlaneGet.json
     */
    /**
     * Sample code: Get packet core control plane.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getPacketCoreControlPlane(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .packetCoreControlPlanes()
            .getByResourceGroupWithResponse("rg1", "TestPacketCoreCP", com.azure.core.util.Context.NONE);
    }
}
