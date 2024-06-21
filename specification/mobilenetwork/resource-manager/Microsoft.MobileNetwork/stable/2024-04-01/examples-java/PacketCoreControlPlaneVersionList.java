
/**
 * Samples for PacketCoreControlPlaneVersions List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/
     * PacketCoreControlPlaneVersionList.json
     */
    /**
     * Sample code: Get supported packet core control plane versions.
     * 
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getSupportedPacketCoreControlPlaneVersions(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.packetCoreControlPlaneVersions().list(com.azure.core.util.Context.NONE);
    }
}
