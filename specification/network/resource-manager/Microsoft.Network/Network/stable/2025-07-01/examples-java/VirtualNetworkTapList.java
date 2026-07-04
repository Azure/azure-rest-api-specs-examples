
/**
 * Samples for VirtualNetworkTaps ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkTapList.json
     */
    /**
     * Sample code: List virtual network taps in resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listVirtualNetworkTapsInResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkTaps().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
