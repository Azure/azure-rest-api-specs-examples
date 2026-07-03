
/**
 * Samples for NetworkManagers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerListAll.json
     */
    /**
     * Sample code: NetworkManagersList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void networkManagersList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkManagers().list(null, null, com.azure.core.util.Context.NONE);
    }
}
