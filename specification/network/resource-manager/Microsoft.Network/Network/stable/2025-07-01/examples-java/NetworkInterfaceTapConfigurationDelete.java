
/**
 * Samples for NetworkInterfaceTapConfigurations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkInterfaceTapConfigurationDelete.json
     */
    /**
     * Sample code: Delete tap configuration.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteTapConfiguration(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkInterfaceTapConfigurations().delete("testrg", "mynic", "tapconfiguration1",
            com.azure.core.util.Context.NONE);
    }
}
