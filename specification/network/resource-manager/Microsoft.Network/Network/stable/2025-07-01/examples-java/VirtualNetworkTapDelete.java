
/**
 * Samples for VirtualNetworkTaps Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkTapDelete.json
     */
    /**
     * Sample code: Delete Virtual Network Tap resource.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteVirtualNetworkTapResource(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkTaps().delete("rg1", "test-vtap", com.azure.core.util.Context.NONE);
    }
}
