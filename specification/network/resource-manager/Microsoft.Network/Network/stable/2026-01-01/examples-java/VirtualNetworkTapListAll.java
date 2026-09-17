
/**
 * Samples for VirtualNetworkTaps List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/VirtualNetworkTapListAll.json
     */
    /**
     * Sample code: List all virtual network taps.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllVirtualNetworkTaps(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkTaps().list(com.azure.core.util.Context.NONE);
    }
}
