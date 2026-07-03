
/**
 * Samples for VirtualHubs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualHubList.json
     */
    /**
     * Sample code: VirtualHubList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void virtualHubList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualHubs().list(com.azure.core.util.Context.NONE);
    }
}
