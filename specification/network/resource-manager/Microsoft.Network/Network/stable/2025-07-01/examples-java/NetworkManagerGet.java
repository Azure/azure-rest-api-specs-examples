
/**
 * Samples for NetworkManagers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerGet.json
     */
    /**
     * Sample code: NetworkManagersGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void networkManagersGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkManagers().getByResourceGroupWithResponse("rg1", "testNetworkManager",
            com.azure.core.util.Context.NONE);
    }
}
