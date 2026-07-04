
/**
 * Samples for NetworkManagers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerDelete.json
     */
    /**
     * Sample code: NetworkManagersDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void networkManagersDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkManagers().delete("rg1", "testNetworkManager", false,
            com.azure.core.util.Context.NONE);
    }
}
