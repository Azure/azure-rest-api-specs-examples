
/**
 * Samples for PacketCoreControlPlaneVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/
     * PacketCoreControlPlaneVersionGet.json
     */
    /**
     * Sample code: Get packet core control plane version.
     * 
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void
        getPacketCoreControlPlaneVersion(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.packetCoreControlPlaneVersions().getWithResponse("2404.0-1", com.azure.core.util.Context.NONE);
    }
}
