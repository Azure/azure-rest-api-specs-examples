import com.azure.core.util.Context;

/** Samples for Sims List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SimListBySubscription.json
     */
    /**
     * Sample code: List sims in a subscription.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listSimsInASubscription(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.sims().list(Context.NONE);
    }
}
