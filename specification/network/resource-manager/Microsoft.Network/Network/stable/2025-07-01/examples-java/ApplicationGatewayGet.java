
/**
 * Samples for ApplicationGateways GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ApplicationGatewayGet.json
     */
    /**
     * Sample code: Get ApplicationGateway.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getApplicationGateway(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationGateways().getByResourceGroupWithResponse("rg1", "appgw",
            com.azure.core.util.Context.NONE);
    }
}
