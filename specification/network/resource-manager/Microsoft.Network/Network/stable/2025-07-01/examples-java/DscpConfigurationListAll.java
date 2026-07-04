
/**
 * Samples for DscpConfiguration List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/DscpConfigurationListAll.json
     */
    /**
     * Sample code: List all network interfaces.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllNetworkInterfaces(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getDscpConfigurations().list(com.azure.core.util.Context.NONE);
    }
}
