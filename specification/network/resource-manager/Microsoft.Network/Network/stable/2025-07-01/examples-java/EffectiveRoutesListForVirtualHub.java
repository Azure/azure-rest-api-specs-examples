
/**
 * Samples for VirtualHubs GetEffectiveVirtualHubRoutes.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/EffectiveRoutesListForVirtualHub.json
     */
    /**
     * Sample code: Effective Routes for the Virtual Hub.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void effectiveRoutesForTheVirtualHub(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualHubs().getEffectiveVirtualHubRoutes("rg1", "virtualHub1", null,
            com.azure.core.util.Context.NONE);
    }
}
