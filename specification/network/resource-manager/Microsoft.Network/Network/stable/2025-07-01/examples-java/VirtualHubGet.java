
/**
 * Samples for VirtualHubs GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualHubGet.json
     */
    /**
     * Sample code: VirtualHubGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void virtualHubGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualHubs().getByResourceGroupWithResponse("rg1", "virtualHub1",
            com.azure.core.util.Context.NONE);
    }
}
