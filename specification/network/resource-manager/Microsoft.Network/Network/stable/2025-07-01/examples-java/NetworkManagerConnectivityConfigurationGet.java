
/**
 * Samples for ConnectivityConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerConnectivityConfigurationGet.json
     */
    /**
     * Sample code: ConnectivityConfigurationsGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void connectivityConfigurationsGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getConnectivityConfigurations().getWithResponse("myResourceGroup", "testNetworkManager",
            "myTestConnectivityConfig", com.azure.core.util.Context.NONE);
    }
}
