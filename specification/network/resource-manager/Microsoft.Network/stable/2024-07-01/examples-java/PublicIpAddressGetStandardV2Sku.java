
/**
 * Samples for PublicIpAddresses GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/
     * PublicIpAddressGetStandardV2Sku.json
     */
    /**
     * Sample code: Get public IP address with StandardV2 sku.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getPublicIPAddressWithStandardV2Sku(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPublicIpAddresses().getByResourceGroupWithResponse("rg1",
            "testDNS-ip", null, com.azure.core.util.Context.NONE);
    }
}
