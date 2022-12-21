import com.azure.core.util.Context;

/** Samples for PacketCoreControlPlaneVersions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/PacketCoreControlPlaneVersionGet.json
     */
    /**
     * Sample code: Get packet core control plane version.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getPacketCoreControlPlaneVersion(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.packetCoreControlPlaneVersions().getWithResponse("PMN-4-11-1", Context.NONE);
    }
}
