
/**
 * Samples for NetworkInterfaceIpConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkInterfaceIPConfigurationList.json
     */
    /**
     * Sample code: NetworkInterfaceIPConfigurationList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void networkInterfaceIPConfigurationList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkInterfaceIpConfigurations().list("testrg", "nic1",
            com.azure.core.util.Context.NONE);
    }
}
