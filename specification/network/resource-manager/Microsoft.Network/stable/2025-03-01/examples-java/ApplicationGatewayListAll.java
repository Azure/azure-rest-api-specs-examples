
/**
 * Samples for ApplicationGateways List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ApplicationGatewayListAll.
     * json
     */
    /**
     * Sample code: Lists all application gateways in a subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listsAllApplicationGatewaysInASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getApplicationGateways().list(com.azure.core.util.Context.NONE);
    }
}
