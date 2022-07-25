import com.azure.core.util.Context;

/** Samples for PacketCoreControlPlanes ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/PacketCoreControlPlaneListByResourceGroup.json
     */
    /**
     * Sample code: List packet core control planes in resource group.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listPacketCoreControlPlanesInResourceGroup(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.packetCoreControlPlanes().listByResourceGroup("rg1", Context.NONE);
    }
}
