/** Samples for MobileNetworks List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-06-01/examples/MobileNetworkListBySubscription.json
     */
    /**
     * Sample code: List mobile networks in a subscription.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listMobileNetworksInASubscription(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.mobileNetworks().list(com.azure.core.util.Context.NONE);
    }
}
