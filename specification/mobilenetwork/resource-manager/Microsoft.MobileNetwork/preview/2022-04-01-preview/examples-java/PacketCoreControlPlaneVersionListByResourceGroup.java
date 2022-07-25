import com.azure.core.util.Context;

/** Samples for PacketCoreControlPlaneVersions ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/PacketCoreControlPlaneVersionListByResourceGroup.json
     */
    /**
     * Sample code: Get supported packet core control plane versions.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getSupportedPacketCoreControlPlaneVersions(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.packetCoreControlPlaneVersions().listByResourceGroup(Context.NONE);
    }
}
