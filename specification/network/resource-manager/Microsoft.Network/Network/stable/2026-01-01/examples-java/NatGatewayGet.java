
/**
 * Samples for NatGateways GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/NatGatewayGet.json
     */
    /**
     * Sample code: Get nat gateway.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getNatGateway(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNatGateways().getByResourceGroupWithResponse("rg1", "test-natGateway", null,
            com.azure.core.util.Context.NONE);
    }
}
