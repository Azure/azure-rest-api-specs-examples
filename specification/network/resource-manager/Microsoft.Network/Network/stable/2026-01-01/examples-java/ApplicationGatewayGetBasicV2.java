
/**
 * Samples for ApplicationGateways GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/ApplicationGatewayGetBasicV2.json
     */
    /**
     * Sample code: Get Basic_v2 Application Gateway.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getBasicV2ApplicationGateway(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationGateways().getByResourceGroupWithResponse("rg1", "appgw",
            com.azure.core.util.Context.NONE);
    }
}
