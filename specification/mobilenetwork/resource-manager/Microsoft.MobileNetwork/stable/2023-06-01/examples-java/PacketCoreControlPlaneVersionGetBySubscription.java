/** Samples for PacketCoreControlPlaneVersions GetBySubscription. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-06-01/examples/PacketCoreControlPlaneVersionGetBySubscription.json
     */
    /**
     * Sample code: Get packet core control plane version by subscription.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getPacketCoreControlPlaneVersionBySubscription(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .packetCoreControlPlaneVersions()
            .getBySubscriptionWithResponse("PMN-4-11-1", com.azure.core.util.Context.NONE);
    }
}
