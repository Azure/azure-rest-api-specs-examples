import com.azure.core.util.Context;

/** Samples for Vendors List. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/VendorListBySubscription.json
     */
    /**
     * Sample code: List all vendor resources in subscription.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void listAllVendorResourcesInSubscription(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.vendors().list(Context.NONE);
    }
}
