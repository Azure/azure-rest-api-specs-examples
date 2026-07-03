
/**
 * Samples for NatGateways GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NatGatewayWithServiceGatewayGet.json
     */
    /**
     * Sample code: Get nat gateway with service gateway.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getNatGatewayWithServiceGateway(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNatGateways().getByResourceGroupWithResponse("rg1", "test-natGateway", null,
            com.azure.core.util.Context.NONE);
    }
}
