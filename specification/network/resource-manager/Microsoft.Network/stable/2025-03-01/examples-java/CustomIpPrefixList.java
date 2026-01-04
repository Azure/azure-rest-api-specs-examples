
/**
 * Samples for CustomIpPrefixes ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/CustomIpPrefixList.json
     */
    /**
     * Sample code: List resource group Custom IP prefixes.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listResourceGroupCustomIPPrefixes(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getCustomIpPrefixes().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
