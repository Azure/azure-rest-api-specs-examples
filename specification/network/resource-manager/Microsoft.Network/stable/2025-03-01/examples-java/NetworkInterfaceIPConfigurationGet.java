
/**
 * Samples for NetworkInterfaceIpConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * NetworkInterfaceIPConfigurationGet.json
     */
    /**
     * Sample code: NetworkInterfaceIPConfigurationGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void networkInterfaceIPConfigurationGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkInterfaceIpConfigurations().getWithResponse("testrg",
            "mynic", "ipconfig1", com.azure.core.util.Context.NONE);
    }
}
