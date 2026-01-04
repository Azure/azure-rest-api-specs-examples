
/**
 * Samples for PublicIpPrefixes GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * PublicIpPrefixGetStandardV2Sku.json
     */
    /**
     * Sample code: Get public IP prefix with StandardV2 sku.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getPublicIPPrefixWithStandardV2Sku(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPublicIpPrefixes().getByResourceGroupWithResponse("rg1",
            "test-ipprefix", null, com.azure.core.util.Context.NONE);
    }
}
