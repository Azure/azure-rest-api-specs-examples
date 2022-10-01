import com.azure.core.util.Context;

/** Samples for Devices List. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/DeviceListBySubscription.json
     */
    /**
     * Sample code: List all devices in a subscription.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void listAllDevicesInASubscription(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.devices().list(Context.NONE);
    }
}
