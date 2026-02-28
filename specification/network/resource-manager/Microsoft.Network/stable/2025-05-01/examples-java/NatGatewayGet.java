
/**
 * Samples for NatGateways GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NatGatewayGet.json
     */
    /**
     * Sample code: Get nat gateway.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getNatGateway(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNatGateways().getByResourceGroupWithResponse("rg1",
            "test-natGateway", null, com.azure.core.util.Context.NONE);
    }
}
