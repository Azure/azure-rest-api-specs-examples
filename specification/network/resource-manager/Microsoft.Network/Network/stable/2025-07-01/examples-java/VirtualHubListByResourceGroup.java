
/**
 * Samples for VirtualHubs ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualHubListByResourceGroup.json
     */
    /**
     * Sample code: VirtualHubListByResourceGroup.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void virtualHubListByResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualHubs().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
