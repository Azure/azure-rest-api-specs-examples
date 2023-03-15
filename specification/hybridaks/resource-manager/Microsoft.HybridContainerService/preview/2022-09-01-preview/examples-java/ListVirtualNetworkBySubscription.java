/** Samples for VirtualNetworksOperation List. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/ListVirtualNetworkBySubscription.json
     */
    /**
     * Sample code: ListVirtualNetworkBySubscription.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void listVirtualNetworkBySubscription(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager.virtualNetworksOperations().list(com.azure.core.util.Context.NONE);
    }
}
