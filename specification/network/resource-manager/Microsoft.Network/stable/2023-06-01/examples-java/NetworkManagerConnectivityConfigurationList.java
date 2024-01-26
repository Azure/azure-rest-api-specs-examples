
/**
 * Samples for ConnectivityConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/
     * NetworkManagerConnectivityConfigurationList.json
     */
    /**
     * Sample code: ConnectivityConfigurationsList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void connectivityConfigurationsList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getConnectivityConfigurations().list("myResourceGroup",
            "testNetworkManager", null, null, com.azure.core.util.Context.NONE);
    }
}
