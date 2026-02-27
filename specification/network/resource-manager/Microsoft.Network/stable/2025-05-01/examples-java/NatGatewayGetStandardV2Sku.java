
/**
 * Samples for NatGateways GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NatGatewayGetStandardV2Sku.
     * json
     */
    /**
     * Sample code: Get nat gateway with StandardV2 Sku.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getNatGatewayWithStandardV2Sku(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNatGateways().getByResourceGroupWithResponse("rg1",
            "test-natGateway", null, com.azure.core.util.Context.NONE);
    }
}
