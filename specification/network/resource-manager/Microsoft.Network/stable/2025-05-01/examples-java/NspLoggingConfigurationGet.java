
/**
 * Samples for NetworkSecurityPerimeterLoggingConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NspLoggingConfigurationGet.
     * json
     */
    /**
     * Sample code: NspLoggingConfigurationGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nspLoggingConfigurationGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeterLoggingConfigurations()
            .getWithResponse("rg1", "nsp1", "instance", com.azure.core.util.Context.NONE);
    }
}
