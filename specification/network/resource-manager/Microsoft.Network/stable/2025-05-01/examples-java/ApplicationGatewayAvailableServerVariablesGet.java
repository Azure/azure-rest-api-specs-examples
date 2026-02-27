
/**
 * Samples for ApplicationGateways ListAvailableServerVariables.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * ApplicationGatewayAvailableServerVariablesGet.json
     */
    /**
     * Sample code: Get Available Server Variables.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAvailableServerVariables(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getApplicationGateways()
            .listAvailableServerVariablesWithResponse(com.azure.core.util.Context.NONE);
    }
}
