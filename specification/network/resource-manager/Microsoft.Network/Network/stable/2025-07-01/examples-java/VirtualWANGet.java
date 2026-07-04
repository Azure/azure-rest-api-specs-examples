
/**
 * Samples for VirtualWans GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualWANGet.json
     */
    /**
     * Sample code: VirtualWANGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void virtualWANGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualWans().getByResourceGroupWithResponse("rg1", "wan1",
            com.azure.core.util.Context.NONE);
    }
}
