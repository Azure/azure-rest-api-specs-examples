
/**
 * Samples for DscpConfiguration GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/DscpConfigurationGet.json
     */
    /**
     * Sample code: Get Dscp Configuration.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDscpConfiguration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getDscpConfigurations().getByResourceGroupWithResponse("rg1",
            "mydscpConfig", com.azure.core.util.Context.NONE);
    }
}
