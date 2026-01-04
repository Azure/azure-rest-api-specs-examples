
/**
 * Samples for NetworkInterfaceTapConfigurations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * NetworkInterfaceTapConfigurationDelete.json
     */
    /**
     * Sample code: Delete tap configuration.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteTapConfiguration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkInterfaceTapConfigurations().delete("testrg", "mynic",
            "tapconfiguration1", com.azure.core.util.Context.NONE);
    }
}
