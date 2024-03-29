import com.azure.core.util.Context;

/** Samples for PacketCoreControlPlanes List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/PacketCoreControlPlaneListBySubscription.json
     */
    /**
     * Sample code: List packet core control planes in a subscription.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listPacketCoreControlPlanesInASubscription(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.packetCoreControlPlanes().list(Context.NONE);
    }
}
