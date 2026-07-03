
/**
 * Samples for NetworkSecurityPerimeters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkSecurityPerimeterGet.json
     */
    /**
     * Sample code: NetworkSecurityPerimeterGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void networkSecurityPerimeterGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeters().getByResourceGroupWithResponse("rg1", "nsp1",
            com.azure.core.util.Context.NONE);
    }
}
