import com.azure.core.util.Context;

/** Samples for VirtualNetworksOperation List. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-05-01-preview/examples/ListVirtualNetworkBySubscription.json
     */
    /**
     * Sample code: ListVirtualNetworkBySubscription.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void listVirtualNetworkBySubscription(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager.virtualNetworksOperations().list(Context.NONE);
    }
}
