
/**
 * Samples for VirtualHubs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualHubDelete.json
     */
    /**
     * Sample code: VirtualHubDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void virtualHubDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualHubs().delete("rg1", "virtualHub1", com.azure.core.util.Context.NONE);
    }
}
