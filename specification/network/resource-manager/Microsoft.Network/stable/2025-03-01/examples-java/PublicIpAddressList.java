
/**
 * Samples for PublicIpAddresses ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/PublicIpAddressList.json
     */
    /**
     * Sample code: List resource group public IP addresses.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listResourceGroupPublicIPAddresses(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPublicIpAddresses().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
