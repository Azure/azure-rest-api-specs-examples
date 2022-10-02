import com.azure.core.util.Context;

/** Samples for NetworkFunctions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/NetworkFunctionListBySubscription.json
     */
    /**
     * Sample code: List all network function resources in subscription.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void listAllNetworkFunctionResourcesInSubscription(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkFunctions().list(Context.NONE);
    }
}
