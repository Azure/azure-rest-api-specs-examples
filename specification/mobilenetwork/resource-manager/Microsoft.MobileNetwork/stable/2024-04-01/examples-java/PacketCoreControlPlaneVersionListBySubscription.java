
/**
 * Samples for PacketCoreControlPlaneVersions ListBySubscription.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/
     * PacketCoreControlPlaneVersionListBySubscription.json
     */
    /**
     * Sample code: Get supported packet core control plane versions by subscription.
     * 
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getSupportedPacketCoreControlPlaneVersionsBySubscription(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.packetCoreControlPlaneVersions().listBySubscription(com.azure.core.util.Context.NONE);
    }
}
