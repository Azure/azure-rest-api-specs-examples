
/**
 * Samples for DscpConfiguration ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/DscpConfigurationList.json
     */
    /**
     * Sample code: Get Dscp Configuration.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDscpConfiguration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getDscpConfigurations().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
