
/**
 * Samples for PacketCoreControlPlaneVersions GetBySubscription.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/
     * PacketCoreControlPlaneVersionGetBySubscription.json
     */
    /**
     * Sample code: Get packet core control plane version by subscription.
     * 
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getPacketCoreControlPlaneVersionBySubscription(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.packetCoreControlPlaneVersions().getBySubscriptionWithResponse("2404.0-1",
            com.azure.core.util.Context.NONE);
    }
}
