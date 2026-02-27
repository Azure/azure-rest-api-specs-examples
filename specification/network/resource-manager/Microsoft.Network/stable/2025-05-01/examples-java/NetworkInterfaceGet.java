
/**
 * Samples for NetworkInterfaces GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkInterfaceGet.json
     */
    /**
     * Sample code: Get network interface.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getNetworkInterface(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkInterfaces().getByResourceGroupWithResponse("rg1",
            "test-nic", null, com.azure.core.util.Context.NONE);
    }
}
