
/**
 * Samples for ApplicationGateways ListAvailableServerVariables.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ApplicationGatewayAvailableServerVariablesGet.json
     */
    /**
     * Sample code: Get Available Server Variables.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getAvailableServerVariables(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationGateways()
            .listAvailableServerVariablesWithResponse(com.azure.core.util.Context.NONE);
    }
}
