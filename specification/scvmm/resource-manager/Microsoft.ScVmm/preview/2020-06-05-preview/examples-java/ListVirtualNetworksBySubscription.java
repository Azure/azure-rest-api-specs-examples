import com.azure.core.util.Context;

/** Samples for VirtualNetworks List. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/ListVirtualNetworksBySubscription.json
     */
    /**
     * Sample code: ListVirtualNetworksBySubscription.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void listVirtualNetworksBySubscription(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualNetworks().list(Context.NONE);
    }
}
