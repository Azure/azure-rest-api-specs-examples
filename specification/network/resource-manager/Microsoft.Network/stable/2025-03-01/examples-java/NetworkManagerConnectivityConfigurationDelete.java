
/**
 * Samples for ConnectivityConfigurations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * NetworkManagerConnectivityConfigurationDelete.json
     */
    /**
     * Sample code: ConnectivityConfigurationsDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void connectivityConfigurationsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getConnectivityConfigurations().delete("myResourceGroup",
            "testNetworkManager", "myTestConnectivityConfig", false, com.azure.core.util.Context.NONE);
    }
}
