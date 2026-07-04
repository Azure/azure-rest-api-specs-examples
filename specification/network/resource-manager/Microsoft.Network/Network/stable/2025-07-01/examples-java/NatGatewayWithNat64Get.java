
/**
 * Samples for NatGateways GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NatGatewayWithNat64Get.json
     */
    /**
     * Sample code: Get nat gateway with nat64.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getNatGatewayWithNat64(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNatGateways().getByResourceGroupWithResponse("rg1", "test-natGateway", null,
            com.azure.core.util.Context.NONE);
    }
}
