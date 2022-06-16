import com.azure.core.util.Context;

/** Samples for MobileNetworks List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/MobileNetworkListBySubscription.json
     */
    /**
     * Sample code: List mobile networks in a subscription.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listMobileNetworksInASubscription(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.mobileNetworks().list(Context.NONE);
    }
}
