
/**
 * Samples for NetworkInterfaceTapConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkInterfaceTapConfigurationList.json
     */
    /**
     * Sample code: List virtual network tap configurations.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listVirtualNetworkTapConfigurations(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkInterfaceTapConfigurations().list("testrg", "mynic",
            com.azure.core.util.Context.NONE);
    }
}
