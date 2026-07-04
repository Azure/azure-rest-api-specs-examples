
/**
 * Samples for ApplicationGateways List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ApplicationGatewayListAll.json
     */
    /**
     * Sample code: Lists all application gateways in a subscription.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listsAllApplicationGatewaysInASubscription(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationGateways().list(com.azure.core.util.Context.NONE);
    }
}
