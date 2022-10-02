import com.azure.core.util.Context;

/** Samples for NetworkFunctionVendors List. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/NetworkFunctionVendorAndSkuListBySubscription.json
     */
    /**
     * Sample code: List vendors and skus.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void listVendorsAndSkus(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkFunctionVendors().list(Context.NONE);
    }
}
