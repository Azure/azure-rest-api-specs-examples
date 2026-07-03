
/**
 * Samples for NetworkInterfaceIpConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkInterfaceIPConfigurationList.json
     */
    /**
     * Sample code: NetworkInterfaceIPConfigurationList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void networkInterfaceIPConfigurationList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkInterfaceIpConfigurations().list("testrg", "nic1",
            com.azure.core.util.Context.NONE);
    }
}
