
/**
 * Samples for ApplicationGateways Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ApplicationGatewayDelete.json
     */
    /**
     * Sample code: Delete ApplicationGateway.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteApplicationGateway(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationGateways().delete("rg1", "appgw", com.azure.core.util.Context.NONE);
    }
}
