
/**
 * Samples for ConnectivityConfigurations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerConnectivityConfigurationDelete.json
     */
    /**
     * Sample code: ConnectivityConfigurationsDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void connectivityConfigurationsDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getConnectivityConfigurations().delete("myResourceGroup", "testNetworkManager",
            "myTestConnectivityConfig", false, com.azure.core.util.Context.NONE);
    }
}
