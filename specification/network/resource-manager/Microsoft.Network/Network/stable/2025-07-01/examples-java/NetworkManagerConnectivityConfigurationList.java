
/**
 * Samples for ConnectivityConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerConnectivityConfigurationList.json
     */
    /**
     * Sample code: ConnectivityConfigurationsList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void connectivityConfigurationsList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getConnectivityConfigurations().list("myResourceGroup", "testNetworkManager", null,
            null, com.azure.core.util.Context.NONE);
    }
}
