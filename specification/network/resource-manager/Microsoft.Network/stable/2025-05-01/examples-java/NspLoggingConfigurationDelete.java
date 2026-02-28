
/**
 * Samples for NetworkSecurityPerimeterLoggingConfigurations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NspLoggingConfigurationDelete
     * .json
     */
    /**
     * Sample code: NspLoggingConfigurationDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nspLoggingConfigurationDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeterLoggingConfigurations()
            .deleteWithResponse("rg1", "nsp1", "instance", com.azure.core.util.Context.NONE);
    }
}
