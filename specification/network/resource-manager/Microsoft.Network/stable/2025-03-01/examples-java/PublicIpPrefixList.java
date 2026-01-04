
/**
 * Samples for PublicIpPrefixes ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/PublicIpPrefixList.json
     */
    /**
     * Sample code: List resource group public IP prefixes.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listResourceGroupPublicIPPrefixes(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPublicIpPrefixes().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
