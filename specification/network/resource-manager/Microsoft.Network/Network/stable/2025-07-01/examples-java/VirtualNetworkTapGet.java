
/**
 * Samples for VirtualNetworkTaps GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkTapGet.json
     */
    /**
     * Sample code: Get Virtual Network Tap.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getVirtualNetworkTap(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkTaps().getByResourceGroupWithResponse("rg1", "testvtap",
            com.azure.core.util.Context.NONE);
    }
}
