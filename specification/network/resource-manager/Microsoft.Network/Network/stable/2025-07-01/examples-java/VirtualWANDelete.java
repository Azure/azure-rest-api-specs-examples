
/**
 * Samples for VirtualWans Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualWANDelete.json
     */
    /**
     * Sample code: VirtualWANDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void virtualWANDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualWans().delete("rg1", "virtualWan1", com.azure.core.util.Context.NONE);
    }
}
