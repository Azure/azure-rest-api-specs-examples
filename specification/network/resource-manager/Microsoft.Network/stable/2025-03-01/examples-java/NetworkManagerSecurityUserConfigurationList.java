
/**
 * Samples for SecurityUserConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * NetworkManagerSecurityUserConfigurationList.json
     */
    /**
     * Sample code: List security user configurations in a network manager.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listSecurityUserConfigurationsInANetworkManager(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getSecurityUserConfigurations().list("rg1", "testNetworkManager",
            null, null, com.azure.core.util.Context.NONE);
    }
}
