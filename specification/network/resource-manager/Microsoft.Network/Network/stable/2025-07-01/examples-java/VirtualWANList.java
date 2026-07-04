
/**
 * Samples for VirtualWans List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualWANList.json
     */
    /**
     * Sample code: VirtualWANList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void virtualWANList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualWans().list(com.azure.core.util.Context.NONE);
    }
}
