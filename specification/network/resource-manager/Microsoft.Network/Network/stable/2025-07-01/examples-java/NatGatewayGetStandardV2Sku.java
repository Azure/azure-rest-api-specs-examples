
/**
 * Samples for NatGateways GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NatGatewayGetStandardV2Sku.json
     */
    /**
     * Sample code: Get nat gateway with StandardV2 Sku.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getNatGatewayWithStandardV2Sku(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNatGateways().getByResourceGroupWithResponse("rg1", "test-natGateway", null,
            com.azure.core.util.Context.NONE);
    }
}
