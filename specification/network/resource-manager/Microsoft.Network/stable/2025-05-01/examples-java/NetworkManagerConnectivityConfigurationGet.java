
/**
 * Samples for ConnectivityConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkManagerConnectivityConfigurationGet.json
     */
    /**
     * Sample code: ConnectivityConfigurationsGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void connectivityConfigurationsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getConnectivityConfigurations().getWithResponse("myResourceGroup",
            "testNetworkManager", "myTestConnectivityConfig", com.azure.core.util.Context.NONE);
    }
}
