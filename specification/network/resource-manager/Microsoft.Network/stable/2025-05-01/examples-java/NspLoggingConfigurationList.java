
/**
 * Samples for NetworkSecurityPerimeterLoggingConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NspLoggingConfigurationList.
     * json
     */
    /**
     * Sample code: NspLoggingConfigurationList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nspLoggingConfigurationList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeterLoggingConfigurations().list("rg1",
            "nsp1", com.azure.core.util.Context.NONE);
    }
}
