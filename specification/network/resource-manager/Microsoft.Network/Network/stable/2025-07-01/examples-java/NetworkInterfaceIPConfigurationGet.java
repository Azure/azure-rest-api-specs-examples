
/**
 * Samples for NetworkInterfaceIpConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkInterfaceIPConfigurationGet.json
     */
    /**
     * Sample code: NetworkInterfaceIPConfigurationGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void networkInterfaceIPConfigurationGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkInterfaceIpConfigurations().getWithResponse("testrg", "mynic", "ipconfig1",
            com.azure.core.util.Context.NONE);
    }
}
